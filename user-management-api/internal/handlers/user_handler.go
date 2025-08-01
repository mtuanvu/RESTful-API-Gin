package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mtuanvu.id.vn/restful-api-gin/internal/dtos"
	"mtuanvu.id.vn/restful-api-gin/internal/models"
	"mtuanvu.id.vn/restful-api-gin/internal/services"
	"mtuanvu.id.vn/restful-api-gin/internal/utils"
	"mtuanvu.id.vn/restful-api-gin/internal/validations"
)

type UserHandler struct {
	service services.UserService
}

type GetUserByUUIDParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

type GetUsersParam struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int    `form:"page" binding:"omitempty,gte=1,lte=100"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewUserHandler(services services.UserService) *UserHandler {
	return &UserHandler{
		service: services,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params GetUsersParam

	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	users, err := uh.service.GetAllUsers(params.Search, params.Page, params.Limit)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userResponse := dtos.MapUsersToDTO(users)
	utils.ResponseSuccess(ctx, http.StatusOK, userResponse)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	createdUser, err := uh.service.CreateUser(user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userResponse := dtos.MapUserToDTO(createdUser)

	utils.ResponseSuccess(ctx, http.StatusCreated, &userResponse)
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	user, err := uh.service.GetUserByUUID(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userResponse := dtos.MapUserToDTO(user)

	utils.ResponseSuccess(ctx, http.StatusOK, userResponse)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	updateUser, err := uh.service.UpdateUser(params.Uuid, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userResponse := dtos.MapUserToDTO(updateUser)

	utils.ResponseSuccess(ctx, http.StatusOK, userResponse)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

}
