package handler

import (
	"net/http"
	"user_microservice/internal/adapter/driver/api/schema"
	"user_microservice/internal/core/domain/port"
	"user_microservice/internal/core/dto"
	"user_microservice/internal/core/use_case"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repository port.IUserRepository
}

func NewUserHandler(repository port.IUserRepository) *UserHandler {
	return &UserHandler{repository: repository}
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var requestBody schema.CreateUserSchema

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.CreateUserDTO{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	createUserUseCase := use_case.NewCreateUserUseCase(uh.repository)

	newUser, err := createUserUseCase.Execute(userDTO)

	if err != nil {
		ctx.Error(err)
		return
	}

	responseBody := schema.UserResponseSchema{
		ID:    newUser.ID,
		Name:  newUser.Name.Value(),
		Email: newUser.Email.Value(),
	}

	ctx.JSON(http.StatusCreated, responseBody)
}

func (uh *UserHandler) FindAllUsers(ctx *gin.Context) {
	findAllUsersUseCase := use_case.NewFindAllUsersUseCase(uh.repository)

	users, err := findAllUsersUseCase.Execute()

	if err != nil {
		ctx.Error(err)
		return
	}

	var responseBody []schema.UserResponseSchema

	for _, user := range users {
		responseBody = append(responseBody, schema.UserResponseSchema{
			ID:    user.ID,
			Name:  user.Name.Value(),
			Email: user.Email.Value(),
		})
	}

	ctx.JSON(http.StatusOK, responseBody)
}

func (uh *UserHandler) FindUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	findUserByIDUseCase := use_case.NewFindUserByIDUseCase(uh.repository)

	user, err := findUserByIDUseCase.Execute(id)

	if err != nil {
		ctx.Error(err)
		return
	}

	responseBody := schema.UserResponseSchema{
		ID:    user.ID,
		Name:  user.Name.Value(),
		Email: user.Email.Value(),
	}

	ctx.JSON(http.StatusOK, responseBody)
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var requestBody schema.UpdateUserSchema

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.UpdateUserDTO{
		ID:    id,
		Name:  requestBody.Name,
		Email: requestBody.Email,
	}

	updateUserUseCase := use_case.NewUpdateUserUseCase(uh.repository)

	updatedUser, err := updateUserUseCase.Execute(userDTO)

	if err != nil {
		ctx.Error(err)
		return
	}

	responseBody := schema.UserResponseSchema{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name.Value(),
		Email: updatedUser.Email.Value(),
	}

	ctx.JSON(http.StatusOK, responseBody)
}

func (uh *UserHandler) UpdateUserPassword(ctx *gin.Context) {
	id := ctx.Param("id")

	var requestBody schema.UpdateUserPasswordSchema

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO := dto.UpdateUserPasswordDTO{
		ID:          id,
		NewPassword: requestBody.NewPassword,
	}

	updateUserPasswordUseCase := use_case.NewUpdateUserPasswordUseCase(uh.repository)

	err := updateUserPasswordUseCase.Execute(userDTO)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	deleteUserUseCase := use_case.NewDeleteUserUseCase(uh.repository)

	err := deleteUserUseCase.Execute(id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (uh *UserHandler) RestoreUser(ctx *gin.Context) {
	id := ctx.Param("id")

	restoreUserUseCase := use_case.NewRestoreUserUseCase(uh.repository)

	err := restoreUserUseCase.Execute(id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
