package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/port"
)

type DeleteUserUseCase struct {
	repository port.IUserRepository
}

func NewDeleteUserUseCase(repository port.IUserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repository: repository,
	}
}

func (u *DeleteUserUseCase) Execute(id string) error {
	err := entity.ValidateID(id)

	if err != nil {
		return err
	}

	err = u.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
