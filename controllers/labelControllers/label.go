package labelcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models/dtos"
	labelrepository "github.com/yornifpaz/back_noteapp/repositories/labelRepository"
)

type ILabelController interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
type LabelController struct {
	repository labelrepository.ILabelRepository
}

// Create implements ILabelController
func (cl *LabelController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := helpers.GetCurrentUserId()
		var body dtos.AddLabel

		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to read body from request",
				"status":  400,
			})
			return

		}
		label, _ := cl.repository.GetByTitle(body.Title, userId)
		if label.ID != "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed label is already",
				"status":  400,
			})
			return
		}
		errLabel := cl.repository.Save(body.Title, userId)
		if errLabel != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create  ",
				"status":  500,
			})
			return
		}
		labels, errLabels := cl.repository.GetAll(userId)
		if errLabels != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load labels ",
				"status":  500,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Created successfully",
			"labels":  labels,
			"status":  200,
		})

	}
}

// Delete implements ILabelController
func (cl *LabelController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		userId := helpers.GetCurrentUserId()
		errDelete := cl.repository.Delete(id)
		if errDelete != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message":   "Failed to delete",
				"isDeleted": false,
				"status":    500,
			})
			return
		}

		labels, errLabels := cl.repository.GetAll(userId)
		if errLabels != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load labels ",
				"status":  500,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Delete  successfully", "isDeleted": true, "status": 200, "labels": labels})

	}
}

// GetAll implements ILabelController
func (cl *LabelController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := helpers.GetCurrentUserId()
		labels, err := cl.repository.GetAll(userId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load labels ",
				"status":  500,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"labels": labels, "status": 200})
	}
}

// Update implements ILabelController
func (cl *LabelController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := helpers.GetCurrentUserId()
		var body dtos.UpdateLabel

		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "Failed to read body from request",
				"isUpdate": false,
				"status":   http.StatusBadRequest,
			})
			return

		}

		// Validate if label title is already
		label, err := cl.repository.GetById(body.Id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message":  "This label no exist ",
				"isUpdate": false,
				"status":   http.StatusBadRequest,
			})
			return
		}

		errUpdate := cl.repository.Update(body.Title, label)
		if errUpdate != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message":  "Failed to update ",
				"isUpdate": false,
				"status":   http.StatusInternalServerError,
			})
			return
		}
		labels, errLabels := cl.repository.GetAll(userId)
		if errLabels != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load labels ",
				"status":  http.StatusInternalServerError,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Updated  successfully",
			"labels":   labels,
			"isUpdate": true,
		})
	}
}

func NewLabelController(repository labelrepository.ILabelRepository) ILabelController {
	return &LabelController{
		repository: repository,
	}
}
