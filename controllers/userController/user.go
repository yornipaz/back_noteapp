package usercontroller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
)

func Register(ctx *gin.Context) {
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Avatar    string
		Password  string
	}
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body from request",
		})
		return

	}
	fmt.Println(body, "IN request body")

	hashPassword, err := helpers.CreatePassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password from",
		})
		return
	}
	var userValidate models.User
	config.DB.First(&userValidate, "email=?", body.Email)
	if userValidate.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "The user already exists in the database",
		})
		return
	}
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: string(hashPassword), Avatar: body.Avatar}

	result := config.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create user in to database",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created user successfully",
	})
}
func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body from request",
		})
		return

	}

	var user models.User
	config.DB.First(&user, "email=?", body.Email)
	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid email address provided or password",
		})
		return
	}

	err := helpers.ValidatePassword(body.Password, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid password provided",
		})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := helpers.CreateJWT(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to sign token",
		})
		return
	}
	ctx.SetSameSite(http.SameSiteDefaultMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	//send response token
	c, err := ctx.Cookie("Authorization")
	fmt.Println("Err : ", err)
	fmt.Println("Cookie : ", c)

	ctx.JSON(http.StatusOK, gin.H{})

}
func Validate(ctx *gin.Context) {
	claims := helpers.GetClaims(helpers.CurrentToken)

	var user models.User
	config.DB.First(&user, claims["sub"])
	if user.ID == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
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
		"message": "Logout successfully user ",
	})

}
