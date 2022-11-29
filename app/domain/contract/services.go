package contract

import (
	"github.com/zakirkun/ehe/app/domain/models"
	"github.com/zakirkun/ehe/app/domain/types"
)

type IAuthService interface {
	SignUpUser(*types.SignUpInput) (*models.DBResponse, error)
	SignInUser(*types.SignInInput) (*models.DBResponse, error)
}

type IUserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
