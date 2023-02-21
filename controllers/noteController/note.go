package notecontroller

import (
	"fmt"
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
	userId := helpers.GetCurrentUserId()
	note := models.Note{
		Title:       body.Title,
		Description: body.Content,
		UserID:      userId,
		Archived:    false,
		Reminder:    time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := config.DB.Create(&note).Error

	if err != nil {
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
	var body dtos.UpdateNote
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	var note models.Note

	errSelect := config.DB.First(&note, "id=?", body.Id).Error
	if errSelect != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "This note no exist ",
			"isUpdate": false,
		})
		return
	}
	var tasks []models.Task
	if len(body.Tasks) > 0 {
		for _, task := range body.Tasks {
			var updateTask = models.Task{
				Id:        task.Id,
				Completed: task.Completed,
				Content:   task.Content,
				NoteRefer: body.Id,
			}
			tasks = append(tasks, updateTask)
		}

	}

	noteUpdate := models.Note{
		UpdatedAt:   time.Now(),
		Reminder:    body.Reminder,
		Title:       body.Title,
		Description: body.Content,
		UserID:      body.UserID,
		Labels:      body.Labels,
		Tasks:       tasks,
	}
	err := config.DB.Model(&note).Updates(noteUpdate).Error

	if err != nil {
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
func UpdateArchived(ctx *gin.Context) {
	var body struct {
		Id       string
		Archived bool
	}
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}
	var note models.Note

	errSelect := config.DB.First(&note, "id=?", body.Id).Error
	if errSelect != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "This note no exist ",
			"isUpdate": false,
		})
		return
	}

	err := config.DB.Model(&note).Updates(map[string]interface{}{"updated_at": time.Now(), "archived": body.Archived}).Error

	if err != nil {
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
	var notes []models.Note
	err := config.DB.Where("user_id = ?", userId).Preload("Tasks").Find(&notes).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get  load ",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notes": notes})
}
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	err := config.DB.Delete(&models.Note{}, "id=?", id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Failed to delete",
			"isDeleted": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete  successfully", "isDeleted": true})

}
func AddTasks(ctx *gin.Context) {
	var body dtos.AddTask
	var Note models.Note
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}

	results := config.DB.First(&Note, "id=?", body.NoteId)

	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "This label no exist ",
			"isUpdate": false,
		})
		return
	}
	task := models.Task{
		Completed: false,
		Content:   body.Task,
		Id:        uuid.NewString(),
		NoteRefer: body.NoteId,
	}
	error := config.DB.Create(&task).Error

	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create  ",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Created successfully",
	})

}
func UpdateTask(ctx *gin.Context) {
	var body dtos.UpdateTask
	var task models.Task
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "Failed to read body from request",
			"isUpdate": false,
		})
		return

	}

	results := config.DB.First(&task, "id=?", body.Id)

	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":  "This task no exist ",
			"isUpdate": false,
		})
		return
	}

	error := config.DB.Model(&task).Updates(map[string]interface{}{"completed": body.Completed, "content": body.Content}).Error

	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update  ",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Update successfully",
		"isUpdate": true,
	})
}
func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	error := config.DB.Delete(&models.Task{}, "id=?", id).Error
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":   "Failed to delete",
			"isDeleted": false,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete  successfully", "isDeleted": true})
}
