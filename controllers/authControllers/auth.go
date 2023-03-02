package authcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userfactory "github.com/yornifpaz/back_noteapp/factory/userFactory"
	"github.com/yornifpaz/back_noteapp/helpers"
	"github.com/yornifpaz/back_noteapp/models/dtos"
	userrepository "github.com/yornifpaz/back_noteapp/repositories/userRepository"
)

type IAuthController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	Validate() gin.HandlerFunc
	Logout() gin.HandlerFunc
}
type AuthController struct {
	repository userrepository.IUserRepository
	factory    userfactory.IUserFactory
}

// Login implements IAuthController
func (cl *AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.LoginUser
		if ctx.Bind(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to read body from request",
				"status":  http.StatusBadRequest,
			})
			return

		}

		user, _ := cl.repository.GetByEmail(body.Email)

		if user.ID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid email address provided ",
				"status":  http.StatusBadRequest,
			})
			return
		}

		err := helpers.ValidatePassword(body.Password, user.Password)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid password provided",
				"status":  http.StatusBadRequest,
			})
			return
		}

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := helpers.CreateJWT(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal error server",
				"status":  http.StatusInternalServerError,
			})
			return
		}
		ctx.SetSameSite(http.SameSiteDefaultMode)
		ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", true, true)
		//send response token

		ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "status": http.StatusOK})
	}
}

// Logout implements IAuthController
func (*AuthController) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Request.Cookie("Authorization")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access ",
				"status":  http.StatusUnauthorized,
			})
			return
		}

		ctx.SetSameSite(http.SameSiteDefaultMode)
		ctx.SetCookie("Authorization", "", -1, "", "", false, true)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Logout successfully",
			"status":  http.StatusOK,
		})
	}
}

// Register implements IAuthController
func (cl *AuthController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.AddUser
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to read body from request",
				"status":  http.StatusBadRequest,
			})
			return

		}

		hashPassword, err := helpers.CreatePassword(body.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal error server  ",
				"status":  http.StatusInternalServerError,
			})
			return
		}
		userValidate, _ := cl.repository.GetByEmail(body.Email)
		if userValidate.Email != "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "The user already exists ",
				"status":  http.StatusBadRequest,
			})
			return
		}
		user := cl.factory.Create(body.FirstName, body.LastName, body.Email, string(hashPassword), body.Avatar)
		errCreate := cl.repository.Save(user)
		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal error server",
				"status":  http.StatusInternalServerError,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Created user successfully",
			"status":  http.StatusOK,
		})
	}
}

// Validate implements IAuthController
func (cl *AuthController) Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := helpers.GetCurrentUserId()
		user, _ := cl.repository.GetById(id)
		if user.ID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access ",
				"status":  http.StatusUnauthorized,
			})
			return

		}
		data := helpers.UserResponse(user)
		ctx.JSON(http.StatusOK, gin.H{
			"user":    data,
			"message": "User Validated successfully ",
			"status":  http.StatusOK,
		})
	}
}

func NewAuthController(repository userrepository.IUserRepository, factory userfactory.IUserFactory) IAuthController {
	return &AuthController{
		repository: repository,
		factory:    factory,
	}
}
