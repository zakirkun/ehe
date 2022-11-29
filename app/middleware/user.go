package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/ehe/app/domain/contract"
	"github.com/zakirkun/ehe/internal/config"
	pkgJwt "github.com/zakirkun/ehe/pkg/jwt"
)

func DeserializeUser(userServices contract.IUserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig("../")

		pkgsJwt := pkgJwt.NewJwt(config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey, config.RefreshTokenPublicKey)
		validateToken, err := pkgsJwt.Validate(cookie)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := userServices.FindUserById(fmt.Sprint(validateToken["userId"]))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()

	}
}
