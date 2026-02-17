package mapper

import (
	"log"
	"user_microservice/internal/adapter/driven/database/model"
	"user_microservice/internal/core/domain/entity"
)

func ToUserModel(user entity.User) model.UserModel {
	return model.UserModel{
		ID:       user.ID,
		Name:     user.Name.Value(),
		Email:    user.Email.Value(),
		Password: user.Password.Value(),
	}
}

func ToUserEntity(userModel model.UserModel) entity.User {
	user, err := entity.NewUserWithoutPassword(
		userModel.ID,
		userModel.Name,
		userModel.Email,
	)

	if err != nil {
		log.Fatalf("Error converting UserModel to User entity: %v", err)
	}

	return *user
}
