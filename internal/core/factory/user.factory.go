package factory

import (
	"user_microservice/internal/adapter/driven/database/repository"
	"user_microservice/internal/common/infra/database"
	"user_microservice/internal/core/domain/port"
)

func NewUserRepository() port.IUserRepository {
	return repository.NewGormUserRepository(database.GetDB())
}
