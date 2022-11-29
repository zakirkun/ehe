package contract

import (
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignUpUser(ctx *gin.Context)
	SignInUser(ctx *gin.Context)
	RefreshAccessToken(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
}

type IUserController interface {
	GetMe(ctx *gin.Context)
}
