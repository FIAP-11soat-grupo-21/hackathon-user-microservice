package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/domain/port"
)

type FindAllUsersUseCase struct {
	repository port.IUserRepository
}

func NewFindAllUsersUseCase(repository port.IUserRepository) *FindAllUsersUseCase {
	return &FindAllUsersUseCase{
		repository: repository,
	}
}

func (f *FindAllUsersUseCase) Execute() ([]entity.User, error) {
	return f.repository.ListAll()
}
