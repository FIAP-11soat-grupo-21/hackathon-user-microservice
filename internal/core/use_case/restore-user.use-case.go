package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/port"
)

type RestoreUserUseCase struct {
	repository  port.IUserRepository
	authService port.IAuthService
}

func NewRestoreUserUseCase(repository port.IUserRepository, authService port.IAuthService) *RestoreUserUseCase {
	return &RestoreUserUseCase{
		repository:  repository,
		authService: authService,
	}
}

func (u *RestoreUserUseCase) Execute(id string) error {
	err := entity.ValidateID(id)

	if err != nil {
		return err
	}

	err = u.repository.Restore(id)

	if err != nil {
		return err
	}

	err = u.authService.RestoreUser(id)

	if err != nil {
		return err
	}

	return nil
}
