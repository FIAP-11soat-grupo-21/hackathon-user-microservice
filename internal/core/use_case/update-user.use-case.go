package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/exception"
	"user_microservice/internal/core/domain/port"
	"user_microservice/internal/core/dto"
)

type UpdateUserUseCase struct {
	repository port.IUserRepository
}

func NewUpdateUserUseCase(repository port.IUserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repository: repository,
	}
}

func (u *UpdateUserUseCase) Execute(userDTO dto.UpdateUserDTO) (entity.User, error) {
	err := entity.ValidateID(userDTO.ID)

	if err != nil {
		return entity.User{}, err
	}

	user, err := u.repository.FindByID(userDTO.ID)

	if err != nil {
		return entity.User{}, err
	}

	if user.IsEmpty() {
		return entity.User{}, &exception.UserNotFoundException{
			Message: "User not found",
		}
	}

	if userDTO.Email != user.Email.Value() {
		err = user.SetEmail(userDTO.Email)

		if err != nil {
			return entity.User{}, err
		}

		emailAlreadyExists, err := u.repository.ExistsByEmail(userDTO.Email)

		if err != nil {
			return entity.User{}, err
		}

		if emailAlreadyExists {
			return entity.User{}, &exception.UserAlreadyExistsException{
				Message: "User with this email already exists",
			}
		}

	}

	if userDTO.Name != user.Name.Value() {
		err = user.SetName(userDTO.Name)

		if err != nil {
			return entity.User{}, err
		}
	}

	err = u.repository.Update(user)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
