package port

import (
	"user_microservice/internal/core/dto"
)

type IAuthService interface {
	RegisterUser(user dto.RegisterUserDTO) error
	UpdateUserEmail(oldEmail, newEmail string) error
	UpdateUserPassword(email, newPassword string) error
	RestoreUser(email string) error
	DeleteUser(email string) error
}
