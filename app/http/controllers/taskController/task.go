package taskcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	taskfactory "github.com/yornifpaz/back_noteapp/app/factory/taskFactory"
	"github.com/yornifpaz/back_noteapp/app/models/dtos"
	noterepository "github.com/yornifpaz/back_noteapp/app/repositories/noteRepository"
	taskrepository "github.com/yornifpaz/back_noteapp/app/repositories/taskRepository"
)

type ITaskController interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
type TaskController struct {
	noteRepository noterepository.INoteRepository
	repository     taskrepository.ITaskRepository
	factory        taskfactory.ITaskFactory
}

// Create implements ITaskController
func (cl *TaskController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.AddTask

		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "Failed to read body from request",
				"isUpdate": false,
			})
			return

		}

		_, errNote := cl.noteRepository.GetById(body.NoteId)

		if errNote != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "This label no exist ",
				"isUpdate": false,
			})
			return
		}
		task := cl.factory.Create(body.Task, body.NoteId)

		err := cl.repository.Save(task)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal error server ",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Created successfully",
		})
	}
}

// Delete implements ITaskController
func (cl *TaskController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := cl.repository.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message":   "Failed to delete",
				"isDeleted": false,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Delete  successfully", "isDeleted": true})
	}
}

// Update implements ITaskController
func (cl *TaskController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.UpdateTask

		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "Failed to read body from request",
				"isUpdate": false,
			})
			return

		}

		task, err := cl.repository.GetById(body.Id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "This task no exist ",
				"isUpdate": false,
			})
			return
		}

		errorUpdate := cl.repository.Update(task, body.Content, body.Completed)

		if errorUpdate != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error ",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Update successfully",
			"isUpdate": true,
		})
	}
}

func NewTaskController(repository taskrepository.ITaskRepository, factory taskfactory.ITaskFactory, noteRepository noterepository.INoteRepository) ITaskController {
	return &TaskController{
		repository:     repository,
		noteRepository: noteRepository,
		factory:        factory,
	}
}
