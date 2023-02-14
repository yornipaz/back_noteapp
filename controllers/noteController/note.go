package notecontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models"
	"github.com/yornifpaz/back_noteapp/models/dtos"
)

func Create(ctx *gin.Context) {
	var body dtos.AddNote
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body from request",
		})
		return

	}

	var tasks []models.Task
	if len(body.Tasks) > 0 {
		for _, task := range body.Tasks {
			var id string = uuid.NewString()
			var newTask = models.Task{
				Id:        id,
				Content:   task.Content,
				Completed: task.Completed,
			}
			tasks = append(tasks, newTask)
		}
	}
	userId := helpers.GetCurrentUserId()
	note := models.Note{
		Title:       body.Title,
		Description: body.Content,
		UserID:      userId,
		Labels:      body.Labels,
		Tasks:       tasks,
		Reminder:    body.Reminder,
		Archived:    body.Archived,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := config.DB.Create(&note)

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
func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var body dtos.UpdateNote

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	var note models.Note
	result := config.DB.First(&note, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "This note no exist ",
			"isUpdate": false,
		})
		return
	}
	noteUpdate := models.Note{
		UpdatedAt:   time.Now(),
		Reminder:    body.Reminder,
		Archived:    body.Archived,
		Title:       body.Title,
		Description: body.Content,
		UserID:      body.UserID,
		Labels:      body.Labels,
		Tasks:       body.Tasks,
	}

	resultUpdate := config.DB.Model(&note).Updates(noteUpdate)
	if resultUpdate.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":  "Failed to update ",
			"isUpdate": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Updated  successfully",
		"note":     note,
		"isUpdate": true,
	})

}
func GetAll(ctx *gin.Context) {
	userId := helpers.GetCurrentUserId()
	var notes []*models.Note
	result := config.DB.Where("user_id = ?", userId).Preload("Tasks").Find(&notes)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get  load ",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Notes": notes})
}
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	result := config.DB.Delete(&models.Note{}, "id=?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Failed to delete",
			"isDeleted": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete  successfully", "isDeleted": true})

}
