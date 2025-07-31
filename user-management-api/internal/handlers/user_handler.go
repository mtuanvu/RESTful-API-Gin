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

func NewUserHandler(services services.UserService) *UserHandler {
	return &UserHandler{
		service: services,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, validations.HandleValidationErrors(err))
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

}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {

}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

}
