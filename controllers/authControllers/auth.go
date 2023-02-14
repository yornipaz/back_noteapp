package authcontrollers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
	"github.com/yornifpaz/back_noteapp/models/dtos"
)

func Register(ctx *gin.Context) {
	var body dtos.AddUser
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body from request",
		})
		return

	}

	hashPassword, err := helpers.CreatePassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal error server  ",
		})
		return
	}
	var userValidate models.User
	config.DB.First(&userValidate, "email=?", body.Email)
	if userValidate.Email != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "The user already exists ",
		})
		return
	}
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: string(hashPassword), Avatar: body.Avatar, CreatedAt: time.Now(), UpdatedAt: time.Now(), Status: "Created"}

	result := config.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user ",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created user successfully",
	})
}
func Login(ctx *gin.Context) {
	var body dtos.LoginUser
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body from request",
		})
		return

	}

	var user models.User
	config.DB.First(&user, "email=?", body.Email)

	if user.ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email address provided ",
		})
		return
	}

	err := helpers.ValidatePassword(body.Password, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password provided",
		})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := helpers.CreateJWT(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to sign token",
		})
		return
	}
	ctx.SetSameSite(http.SameSiteDefaultMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", true, true)
	//send response token

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}
func Validate(ctx *gin.Context) {
	id := helpers.GetCurrentUserId()

	var user models.User
	config.DB.First(&user, "id=?", id)
	if user.ID == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	data := map[string]interface{}{
		"id":         user.ID,
		"avatar":     user.Avatar,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"status":     user.Status,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": data,
	})

}
func Logout(ctx *gin.Context) {

	cookie, err := ctx.Request.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	cookie.Name = "delete_token"
	cookie.Value = "Unset"
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(1, 0)
	ctx.SetSameSite(http.SameSiteDefaultMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully  ",
	})

}
