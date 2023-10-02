package notecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	notefactory "github.com/yornifpaz/back_noteapp/app/factory/noteFactory"
	"github.com/yornifpaz/back_noteapp/app/helpers"
	"github.com/yornifpaz/back_noteapp/app/models/dtos"
	noterepository "github.com/yornifpaz/back_noteapp/app/repositories/noteRepository"
)

type INoteController interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	UpdateArchived() gin.HandlerFunc
}
type NoteController struct {
	repository noterepository.INoteRepository
	factory    notefactory.INoteFactory
}

// Create implements INoteController
func (cl *NoteController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.AddNote
		var userId string
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to read body from request",
			})
			return
		}
		userId = helpers.GetCurrentUserId(ctx)
		note := cl.factory.Create(body.Title, userId, body.Content)
		err := cl.repository.Save(note)
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
}

// Delete implements INoteController
func (cl *NoteController) Delete() gin.HandlerFunc {
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

// GetAll implements INoteController
func (cl *NoteController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := helpers.GetCurrentUserId(ctx)
		notes, err := cl.repository.GetAll(userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get  load ",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"notes": notes})
	}
}

// Update implements INoteController
func (cl *NoteController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.UpdateNote
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "Failed to read body from request",
				"isUpdate": false,
			})
			return

		}

		note, errSelect := cl.repository.GetById(body.Id)
		if errSelect != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "This note no exist ",
				"isUpdate": false,
			})
			return
		}
		tasks := cl.factory.UpdateTask(body.Tasks, body.Id)
		noteFactory := cl.factory.Update(body.Title, body.UserID, body.Content, body.Labels, tasks, body.Reminder)

		noteUpdate, err := cl.repository.Update(note, noteFactory)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message":  "Failed to update ",
				"isUpdate": false,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Updated  successfully",
			"note":     noteUpdate,
			"isUpdate": true,
		})
	}
}

// UpdateArchived implements INoteController
func (cl *NoteController) UpdateArchived() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		note, errSelect := cl.repository.GetById(body.Id)
		if errSelect != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "This note no exist ",
				"isUpdate": false,
			})
			return
		}
		noteArchivedFactory := cl.factory.UpdateArchived(body.Archived)
		noteUpdate, err := cl.repository.UpdateArchived(note, noteArchivedFactory)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message":  "Failed to update ",
				"isUpdate": false,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Updated  successfully",
			"note":     noteUpdate,
			"isUpdate": true,
		})
	}
}

func NewNoteController(repository noterepository.INoteRepository, factory notefactory.INoteFactory) INoteController {
	return &NoteController{repository: repository, factory: factory}
}
