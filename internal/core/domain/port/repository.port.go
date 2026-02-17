package port

import "user_microservice/internal/core/domain/entity"

type IUserRepository interface {
	Insert(user entity.User) error
	ListAll() ([]entity.User, error)
	FindByID(id string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	ExistsByEmail(email string) (bool, error)
	Update(user entity.User) error
	Delete(id string) error
	Restore(id string) error
}
