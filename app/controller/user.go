package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/ehe/app/domain/contract"
	"github.com/zakirkun/ehe/app/domain/models"
	"github.com/zakirkun/ehe/app/helper"
)

type userControllerContext struct {
	userService contract.IUserService
}

func NewUserController(userService contract.IUserService) contract.IUserController {
	return &userControllerContext{userService: userService}
}

func (c *userControllerContext) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": helper.FilteredResponse(currentUser)}})
}
