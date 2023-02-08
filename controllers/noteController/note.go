package notecontroller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
)

func Create(ctx *gin.Context) {
	var body struct {
		Title       string
		Description string
		UserID      string
		Tags        pq.StringArray
	}
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body from request",
		})
		return

	}

	note := models.Note{
		Title:       body.Title,
		Description: body.Description,
		UserID:      body.UserID,
		Tags:        body.Tags,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := config.DB.Create(&note)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create note in to database",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created note successfully",
	})

}
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var body struct {
		Title       string
		Description string
		UserID      string
		Tags        pq.StringArray
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body from request",
		})
		return

	}
	var note models.Note
	result := config.DB.First(&note, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "This note no exist in to database",
		})
		return
	}
	noteUpdate := models.Note{Title: body.Title,
		Description: body.Description,
		UserID:      body.UserID,
		Tags:        body.Tags,
		UpdatedAt:   time.Now(),
	}

	resultUpdate := config.DB.Model(&note).Updates(noteUpdate)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to update note in to database",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated note successfully",
		"note":    note,
	})

}
func GetAll(ctx *gin.Context) {
	claims := helpers.GetClaims(helpers.CurrentToken)
	var notes []*models.Note
	result := config.DB.Where("user_id = ?", claims["sub"]).Find(&notes)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error)
	}
	fmt.Println(notes)
	ctx.JSON(http.StatusOK, gin.H{"Notes": notes})
}
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	result := config.DB.Delete(&models.Note{}, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to delete note in to database",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete note successfully"})

}
