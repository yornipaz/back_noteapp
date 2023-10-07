package usercontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yornifpaz/back_noteapp/app/exceptions"
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
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to read body from request",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return

		}
		user, errUser := cl.repository.GetById(id)

		if errUser != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "This user no exist ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}
		userFactory := cl.factory.Update(body.FirstName, body.LastName, body.Email)

		userUpdate, err := cl.repository.Update(user, userFactory)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to update user ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}
		data := helpers.UserResponse(userUpdate)
		ctx.JSON(http.StatusOK, models.APIResponse{
			Message: "resource updated successfully",
			Status:  http.StatusOK,
			Data:    data,
		})
	}
}

// UpdateAvatar implements IUserController
func (cl *UserController) UpdateAvatar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer exceptions.PanicHandler(ctx)
		log.Info("start update avatar")

		formfile, _, err := ctx.Request.FormFile("avatar")
		if err != nil {
			log.Error("Happened error when mapping request from FE. Error", err)
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Select a image to upload ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}
		id := helpers.GetCurrentUserId(ctx)
		user, errUser := cl.repository.GetById(id)
		if errUser != nil {

			log.Error("Happened error when mapping request from FE. Error", errUser)

			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "This user no exist",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}
		uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
		if err != nil {
			log.Error("Failed to upload avatar to cloud storage : ", err)
			ctx.JSON(
				http.StatusInternalServerError,
				models.APIResponse{
					Message: "Internal Error server",
					Status:  http.StatusInternalServerError,
					Data:    nil,
				},
			)
			return
		}
		userFactory := models.User{
			Avatar:    uploadUrl,
			UpdatedAt: time.Now(),
		}
		userUpdate, errUpdate := cl.repository.Update(user, userFactory)
		if errUpdate != nil {
			log.Error("Failed to update avatar : ", errUpdate)
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Failed to update avatar ",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}

		data := helpers.UserResponse(userUpdate)
		ctx.JSON(
			http.StatusOK, models.APIResponse{
				Message: "success",
				Status:  http.StatusOK,
				Data:    data,
			},
		)
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
			Status:    models.UserStatus(body.Status),
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
