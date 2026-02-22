package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/exception"
	"user_microservice/internal/core/domain/port"
	"user_microservice/internal/core/dto"
)

type UpdateUserPasswordUseCase struct {
	repository  port.IUserRepository
	authService port.IAuthService
}

func NewUpdateUserPasswordUseCase(repository port.IUserRepository, authService port.IAuthService) *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		repository:  repository,
		authService: authService,
	}
}

func (u *UpdateUserPasswordUseCase) Execute(dto dto.UpdateUserPasswordDTO) error {
	err := entity.ValidateID(dto.ID)

	if err != nil {
		return err
	}

	user, err := u.repository.FindByID(dto.ID)

	if err != nil {
		return err
	}

	if user.IsEmpty() {
		return &exception.UserNotFoundException{
			Message: "User not found",
		}
	}

	err = user.SetPassword(dto.NewPassword)

	if err != nil {
		return err
	}

	err = u.repository.Update(user)

	if err != nil {
		return err
	}

	err = u.authService.UpdateUserPassword(user.Email.Value(), dto.NewPassword)

	if err != nil {
		return err
	}

	return nil
}
