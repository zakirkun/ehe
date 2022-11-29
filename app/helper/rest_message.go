package helper

import "github.com/zakirkun/ehe/app/domain/models"

func FilteredResponse(user *models.DBResponse) models.UserResponse {
	return models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
