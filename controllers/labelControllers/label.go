package labelcontrollers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
	"github.com/yornifpaz/back_noteapp/models/dtos"
)

func Create(ctx *gin.Context) {
	var body dtos.AddLabel
	var Label models.Label
	user_id := helpers.GetCurrentUserId()
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body from request",
		})
		return

	}
	// Validate if label title is already
	results := config.DB.Where("user_id=?", user_id).Where("title=?", body.Title).First(&Label)

	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed label is already",
		})
		return
	}
	newLabel := models.Label{Title: body.Title, UserID: user_id, UpdatedAt: time.Now(), CreatedAt: time.Now()}

	result := config.DB.Create(&newLabel)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create  ",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created successfully",
	})

}
func GetAll(ctx *gin.Context) {
	var labels []models.Label
	user_id := helpers.GetCurrentUserId()
	result := config.DB.Where("user_id = ?", user_id).Find(&labels)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to load labels ",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"labels": labels})
}
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var body dtos.UpdateLabel
	var label models.Label
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	// Validate if label title is already
	results := config.DB.First(&label, "id=?", id)

	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "This note no exist ",
			"isUpdate": false,
		})
		return
	}
	updateLabel := models.Label{Title: body.Title, UpdatedAt: time.Now()}

	resultUpdate := config.DB.Model(&label).Updates(updateLabel)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Failed to update ",
			"isUpdate": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Updated  successfully",
		"label":    label,
		"isUpdate": true,
	})
}
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	result := config.DB.Delete(&models.Label{}, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Failed to delete",
			"isDeleted": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete  successfully", "isDeleted": true})
}
