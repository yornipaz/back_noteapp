package usercontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
	"github.com/yornifpaz/back_noteapp/models/dtos"
	"github.com/yornifpaz/back_noteapp/services"
)

func Update(ctx *gin.Context) {

	id := helpers.GetCurrentUserId()

	var body dtos.UpdateUser
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":    "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	var user models.User
	result := config.DB.First(&user, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":    "This user no exist ",
			"isUpdate": false,
		})
		return
	}
	userUpdate := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		UpdatedAt: time.Now(),
	}

	resultUpdate := config.DB.Model(&user).Updates(userUpdate)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":    "Failed to update user ",
			"isUpdate": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Updated  successfully",
		"user":     user,
		"isUpdate": true,
	})
}
func UpdateAvatar(ctx *gin.Context) {

	formfile, _, err := ctx.Request.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Select a image to upload ", "isUpdate": false})
		return
	}
	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Error uploading  image", "isUpdate": false},
		)
		return
	}

	id := helpers.GetCurrentUserId()
	var user models.User
	result := config.DB.First(&user, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":    "This user no exist ",
			"isUpdate": false,
		})
		return
	}
	resultUpdate := config.DB.Model(&user).Update("avatar", uploadUrl)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":    "Failed to update  avatar ",
			"isUpdate": false,
		})
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"statusCode": http.StatusOK,
			"message":    "success",
			"user":       user,
			"isUpdate":   true,
		})

}
func UpdateStatus(ctx *gin.Context) {

	id := helpers.GetCurrentUserId()

	var body struct {
		Status string
	}
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":    "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	var user models.User
	result := config.DB.First(&user, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "This user no exist ",
		})
		return
	}
	userUpdate := models.User{
		Status:    body.Status,
		UpdatedAt: time.Now(),
	}

	resultUpdate := config.DB.Model(&user).Updates(userUpdate)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":    "Failed to update user",
			"isUpdate": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Updated successfully",
		"user":     user,
		"isUpdate": true,
	})
}
func UpdatePassword(ctx *gin.Context) {
	id := helpers.GetCurrentUserId()

	var body struct {
		Password string
	}
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error":    "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	var user models.User
	result := config.DB.First(&user, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "This user no exist ",
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
	userUpdate := models.User{
		Password:  string(hashPassword),
		UpdatedAt: time.Now(),
	}

	resultUpdate := config.DB.Model(&user).Updates(userUpdate)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error":    "Failed to update user",
			"isUpdate": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Updated successfully",
		"user":     user,
		"isUpdate": true,
	})
}
