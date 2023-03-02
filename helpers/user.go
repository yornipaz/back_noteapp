package helpers

import "github.com/yornifpaz/back_noteapp/models"

func UserResponse(user models.User) (data map[string]interface{}) {
	data = map[string]interface{}{
		"id":         user.ID,
		"avatar":     user.Avatar,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"status":     user.Status,
	}
	return
}
