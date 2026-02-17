package use_case

import (
	"user_microservice/internal/common/pkg/identity"
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/exception"
	"user_microservice/internal/core/domain/port"
	"user_microservice/internal/core/dto"
)

type CreateUserUseCase struct {
	repository port.IUserRepository
}

func NewCreateUserUseCase(repository port.IUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repository: repository,
	}
}

func (c *CreateUserUseCase) Execute(userDTO dto.CreateUserDTO) (entity.User, error) {
	emailAlreadyExists, err := c.repository.ExistsByEmail(userDTO.Email)

	if err != nil {
		return entity.User{}, err
	}

	if emailAlreadyExists {
		return entity.User{}, &exception.UserAlreadyExistsException{
			Message: "User with this email already exists",
		}
	}

	id := identity.NewUUIDV7()

	user, err := entity.NewUser(
		id,
		userDTO.Name,
		userDTO.Email,
		userDTO.Password,
	)

	if err != nil {
		return entity.User{}, err
	}

	err = c.repository.Insert(*user)

	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}
