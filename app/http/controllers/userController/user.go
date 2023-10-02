package usercontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	userfactory "github.com/yornifpaz/back_noteapp/app/factory/userFactory"
	"github.com/yornifpaz/back_noteapp/app/helpers"
	"github.com/yornifpaz/back_noteapp/app/models"
	"github.com/yornifpaz/back_noteapp/app/models/dtos"
	userrepository "github.com/yornifpaz/back_noteapp/app/repositories/userRepository"

	"github.com/yornifpaz/back_noteapp/services"
)

type IUserController interface {
	UpdatePassword() gin.HandlerFunc
	UpdateStatus() gin.HandlerFunc
	Update() gin.HandlerFunc
	UpdateAvatar() gin.HandlerFunc
}
type UserController struct {
	repository userrepository.IUserRepository
	factory    userfactory.IUserFactory
}

// Update implements IUserController
func (cl *UserController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := helpers.GetCurrentUserId(ctx)

		var body dtos.UpdateUser
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":    "Failed to read body from request",
				"isUpdate": false,
			})
			return

		}
		user, errUser := cl.repository.GetById(id)

		if errUser != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":    "This user no exist ",
				"isUpdate": false,
			})
			return
		}
		userFactory := cl.factory.Update(body.FirstName, body.LastName, body.Email)

		userUpdate, err := cl.repository.Update(user, userFactory)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":    "Failed to update user ",
				"isUpdate": false,
			})
			return
		}
		data := helpers.UserResponse(userUpdate)
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Updated  successfully",
			"user":     data,
			"isUpdate": true,
		})
	}
}

// UpdateAvatar implements IUserController
func (cl *UserController) UpdateAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		formfile, _, err := ctx.Request.FormFile("avatar")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Select a image to upload ", "isUpdate": false})
			return
		}
		id := helpers.GetCurrentUserId(ctx)
		user, errUser := cl.repository.GetById(id)
		if errUser != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error":    "This user no exist ",
				"isUpdate": false,
			})
			return
		}
		uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Internal Error server", "isUpdate": false},
			)
			return
		}
		userFactory := models.User{
			Avatar:    uploadUrl,
			UpdatedAt: time.Now(),
		}
		userUpdate, errUpdate := cl.repository.Update(user, userFactory)
		if errUpdate != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error":    "Failed to update  avatar ",
				"isUpdate": false,
			})
			return
		}
		data := helpers.UserResponse(userUpdate)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status":   http.StatusOK,
				"message":  "success",
				"user":     data,
				"isUpdate": true,
			})
	}
}

// UpdatePassword implements IUserController
func (cl *UserController) UpdatePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := helpers.GetCurrentUserId(ctx)

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
		user, errUser := cl.repository.GetById(id)
		if errUser != nil {
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
		userFactory := models.User{
			Password:  string(hashPassword),
			UpdatedAt: time.Now(),
		}

		userUpdate, err := cl.repository.Update(user, userFactory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error":    "Internal error Server ",
				"isUpdate": false,
			})
			return
		}

		data := helpers.UserResponse(userUpdate)
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Updated  successfully",
			"user":     data,
			"isUpdate": true,
		})
	}
}

// UpdateStatus implements IUserController
func (cl *UserController) UpdateStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := helpers.GetCurrentUserId(ctx)

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
		user, errUser := cl.repository.GetById(id)
		if errUser != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": "This user no exist ",
			})
			return
		}
		userFactory := models.User{
			Status:    body.Status,
			UpdatedAt: time.Now(),
		}

		userUpdate, err := cl.repository.Update(user, userFactory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error":    "Internal error Server ",
				"isUpdate": false,
			})
			return
		}

		data := helpers.UserResponse(userUpdate)
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Updated  successfully",
			"user":     data,
			"isUpdate": true,
		})
	}
}

func NewUserController(repository userrepository.IUserRepository, factory userfactory.IUserFactory) IUserController {
	return &UserController{
		repository: repository,
		factory:    factory,
	}
}
