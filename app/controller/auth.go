package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/ehe/app/domain/contract"
	"github.com/zakirkun/ehe/app/domain/types"
	"github.com/zakirkun/ehe/app/helper"
	"github.com/zakirkun/ehe/internal/config"
	pkgJwt "github.com/zakirkun/ehe/pkg/jwt"
	"go.mongodb.org/mongo-driver/mongo"
)

type authControllerContext struct {
	authServices contract.IAuthService
	userServices contract.IUserService
}

func NewAuthController(authServices contract.IAuthService, userServices contract.IUserService) contract.IAuthController {
	return &authControllerContext{authServices: authServices, userServices: userServices}
}

func (c *authControllerContext) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (c *authControllerContext) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := config.LoadConfig("../")

	pkgsJwt := pkgJwt.NewJwt(config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey, config.RefreshTokenPublicKey)
	validateToken, err := pkgsJwt.Validate(cookie)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := c.userServices.FindUserById(fmt.Sprint(validateToken["userId"]))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	// time
	now := time.Now().UTC()

	// claims data
	claims := make(map[string]interface{})
	claims["userId"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = now.Add(config.AccessTokenExpiresIn).Unix()

	// access token
	pkgsJwt = pkgJwt.NewJwt(config.AccessTokenExpiresIn, config.AccessTokenPrivateKey, config.AccessTokenPublicKey)

	accessToken, err := pkgsJwt.Generate(claims)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (c *authControllerContext) SignUpUser(ctx *gin.Context) {
	var input *types.SignUpInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if input.Password != input.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	newUser, err := c.authServices.SignUpUser(input)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": helper.FilteredResponse(newUser)}})
}

func (c *authControllerContext) SignInUser(ctx *gin.Context) {

	var input *types.SignInInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := c.userServices.FindUserByEmail(input.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := helper.VerifyPassword(user.Password, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := config.LoadConfig("../")

	// time
	now := time.Now().UTC()

	// claims data
	claims := make(map[string]interface{})
	claims["userId"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = now.Add(config.AccessTokenExpiresIn).Unix()

	// access token
	pkgsJwt := pkgJwt.NewJwt(config.AccessTokenExpiresIn, config.AccessTokenPrivateKey, config.AccessTokenPublicKey)

	accessToken, err := pkgsJwt.Generate(claims)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// refresh token
	pkgsJwt = pkgJwt.NewJwt(config.RefreshTokenExpiresIn, config.RefreshTokenPrivateKey, config.RefreshTokenPublicKey)
	refreshToken, err := pkgsJwt.Generate(claims)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
