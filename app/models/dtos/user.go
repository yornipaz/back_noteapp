package dtos

type AddUser struct {
	FirstName string
	LastName  string
	Email     string
	Avatar    string
	Password  string
}

type LoginUser struct {
	Email    string
	Password string
}
type UpdateUser struct {
	FirstName string
	LastName  string
	Email     string
}
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
type ResetPasswordRequest struct {
	NewPassword     string `json:"newPassword" binding:"required,min=8"`                   // Suponiendo un requisito mínimo de longitud de contraseña
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=NewPassword"` // Verificación de contraseña
}
