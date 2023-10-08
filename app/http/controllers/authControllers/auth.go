package authcontrollers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	constants "github.com/yornifpaz/back_noteapp/app/constant"
	userfactory "github.com/yornifpaz/back_noteapp/app/factory/userFactory"
	"github.com/yornifpaz/back_noteapp/app/helpers"
	"github.com/yornifpaz/back_noteapp/app/models"
	"github.com/yornifpaz/back_noteapp/app/models/dtos"
	authrepository "github.com/yornifpaz/back_noteapp/app/repositories/authRepository"
	userrepository "github.com/yornifpaz/back_noteapp/app/repositories/userRepository"
	"github.com/yornifpaz/back_noteapp/services"
)

type IAuthController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	Validate() gin.HandlerFunc
	Logout() gin.HandlerFunc
	ForgotPassword() gin.HandlerFunc
	ResetPassword() gin.HandlerFunc
}
type AuthController struct {
	repository     userrepository.IUserRepository
	factory        userfactory.IUserFactory
	authRepository authrepository.IAuthRepository
	emailService   services.IEmailService
}

// ForgotPassword implements IAuthController.
func (cl *AuthController) ForgotPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.ForgotPasswordRequest
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to read body from request",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		// Verificar si el usuario existe por su correo electrónico
		user, err := cl.repository.GetByEmail(body.Email)
		if err != nil || user.ID == "" {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Invalid email address provided",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		// Generar un token temporal para restablecimiento de contraseña

		resetToken, err := cl.authRepository.CreateJWT(user.ID, 1, os.Getenv("SECRET_KEY"), nil)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Internal error server",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}

		// Aquí normalmente enviarías el token al usuario (por ejemplo, por correo electrónico)
		data := models.RecoveryEmailData{
			Username:  user.FirstName + " " + user.LastName,
			ResetLink: os.Getenv("FRONTEND_URL") + "/reset-password/" + resetToken,
		}

		message, errEmail := cl.emailService.SendEmail(user.Email, constants.HTML, constants.ForgotPassword, data)

		if errEmail != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: message,
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}
		// En este ejemplo, simplemente lo incluimos en la respuesta para propósitos demostrativos.
		ctx.JSON(http.StatusOK, models.APIResponse{
			Message: "Reset token generated successfully : " + message,
			Status:  http.StatusOK,
			Data: map[string]interface{}{
				"resetToken": resetToken,
			},
		})
	}
}

// ResetPassword implements IAuthController.
func (cl *AuthController) ResetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el resetToken de la URL
		resetToken := ctx.Param("resetToken")
		var body dtos.ResetPasswordRequest
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to read body from request",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		// Verificar la autenticidad del token temporal
		id, err := cl.authRepository.ValidateToken(resetToken)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Invalid reset token",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		// Verificar si el usuario asociado al token existe
		user, err := cl.repository.GetById(id)
		if err != nil || user.ID == "" {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Invalid user associated with the reset token",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		// Restablecer la contraseña del usuario
		hashedPassword, err := cl.authRepository.CreatePassword(body.NewPassword)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Internal error server",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}

		// Actualizar la contraseña del usuario en la base de datos
		userFactory := models.User{
			Password: string(hashedPassword),
		}
		_, err = cl.repository.Update(user, userFactory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Internal error server",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}

		// Aquí podrías enviar una respuesta exitosa o redirigir al usuario a la página de inicio de sesión, etc.
		ctx.JSON(http.StatusNoContent, models.APIResponse{
			Message: "Password reset successfully",
			Status:  http.StatusOK,
			Data:    nil,
		})
	}
}

// Login implements IAuthController
func (cl *AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.LoginUser
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to read body from request",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return

		}

		user, _ := cl.repository.GetByEmail(body.Email)

		if user.ID == "" {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Invalid email address provided ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		err := cl.authRepository.ValidatePassword(body.Password, user.Password)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Invalid password provided ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}

		// Sign and get the complete encoded token as a string using the secret
		payload := map[string]interface{}{
			"roles": user.Roles,
		}
		const expiredToken int64 = 24
		var secretKey string = os.Getenv("SECRET_KEY")
		tokenString, err := cl.authRepository.CreateJWT(user.ID, expiredToken, secretKey, payload)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Internal error server",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}

		//send response token

		ctx.JSON(http.StatusOK, models.APIResponse{
			Message: "Login successful",
			Status:  http.StatusOK,
			Data: map[string]interface{}{
				"token": tokenString,
			}},
		)
	}
}

// Logout implements IAuthController
func (cl *AuthController) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := helpers.GetCurrentUserId(ctx)
		user, err := cl.repository.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to logout ",
				Status:  http.StatusBadRequest,
				Data: map[string]interface{}{
					"isLogout": false,
				},
			})
			return
		}

		userFactory := models.User{
			LogoutAt: time.Now(),
		}
		_, errUpdate := cl.repository.Update(user, userFactory)
		if errUpdate != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Failed to logout ",
				Status:  http.StatusBadRequest,
				Data: map[string]interface{}{
					"isLogout": false,
				},
			})
			return
		}

		ctx.JSON(http.StatusOK,
			models.APIResponse{
				Message: "Logout successfully",
				Status:  http.StatusBadRequest,
				Data: map[string]interface{}{
					"isLogout": true,
				},
			},
		)
	}
}

// Register implements IAuthController
func (cl *AuthController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dtos.AddUser
		if ctx.BindJSON(&body) != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to read body from request",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return

		}

		hashPassword, err := cl.authRepository.CreatePassword(body.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Internal error server",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}
		userValidate, _ := cl.repository.GetByEmail(body.Email)
		if userValidate.Email != "" {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "The user already exists ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}
		user := cl.factory.Create(body.FirstName, body.LastName, body.Email, string(hashPassword), body.Avatar)
		errCreate := cl.repository.Save(user)
		if errCreate != nil {
			ctx.JSON(http.StatusInternalServerError, models.APIResponse{
				Message: "Internal error server",
				Status:  http.StatusInternalServerError,
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusCreated, models.APIResponse{
			Message: "Created user successfully",
			Status:  http.StatusCreated,
			Data:    nil,
		},
		)
	}
}

// Validate implements IAuthController
func (cl *AuthController) Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := helpers.GetCurrentUserId(ctx)
		user, err := cl.repository.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.APIResponse{
				Message: "Failed to user validated  ",
				Status:  http.StatusBadRequest,
				Data:    nil,
			})
			return
		}
		data := helpers.UserResponse(user)
		ctx.JSON(http.StatusOK, models.APIResponse{
			Message: "User Validated successfully",
			Status:  http.StatusOK,
			Data: map[string]interface{}{
				"user": data,
			},
		})
	}
}

func NewAuthController(repository userrepository.IUserRepository, factory userfactory.IUserFactory, authRepository authrepository.IAuthRepository, emailService services.IEmailService) IAuthController {
	return &AuthController{
		repository:     repository,
		factory:        factory,
		emailService:   emailService,
		authRepository: authRepository,
	}
}
