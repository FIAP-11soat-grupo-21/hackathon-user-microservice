package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/exception"
	"user_microservice/internal/core/domain/port"
)

type FindUserByIDUseCase struct {
	repository port.IUserRepository
}

func NewFindUserByIDUseCase(repository port.IUserRepository) *FindUserByIDUseCase {
	return &FindUserByIDUseCase{
		repository: repository,
	}
}

func (u *FindUserByIDUseCase) Execute(id string) (entity.User, error) {
	err := entity.ValidateID(id)

	if err != nil {
		return entity.User{}, err
	}

	user, err := u.repository.FindByID(id)

	if err != nil {
		return entity.User{}, err
	}

	if user.IsEmpty() {
		return entity.User{}, &exception.UserNotFoundException{
			Message: "User not found",
		}
	}

	return user, nil
}
