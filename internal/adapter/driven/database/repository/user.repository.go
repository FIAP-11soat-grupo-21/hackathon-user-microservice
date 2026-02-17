package repository

import (
	"errors"
	"user_microservice/internal/adapter/driven/database/mapper"
	"user_microservice/internal/adapter/driven/database/model"
	"user_microservice/internal/core/domain/entity"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(instance *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: instance}
}

func (r *GormUserRepository) Insert(user entity.User) error {
	userModel := mapper.ToUserModel(user)
	return r.db.Create(&userModel).Error
}

func (r *GormUserRepository) ListAll() ([]entity.User, error) {
	var userModels []model.UserModel

	err := r.db.Find(&userModels).Error
	if err != nil {
		return nil, err
	}

	users := make([]entity.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = mapper.ToUserEntity(userModel)
	}

	return users, nil
}

func (r *GormUserRepository) FindByID(id string) (entity.User, error) {
	var userModel model.UserModel

	err := r.db.First(&userModel, "id = ?", id).Error
	if err != nil {
		return entity.User{}, err
	}

	return mapper.ToUserEntity(userModel), nil
}

func (r *GormUserRepository) FindByEmail(email string) (entity.User, error) {
	var userModel model.UserModel

	err := r.db.First(&userModel, "email = ?", email).Error
	if err != nil {
		return entity.User{}, err
	}

	return mapper.ToUserEntity(userModel), nil
}

func (r *GormUserRepository) ExistsByEmail(email string) (bool, error) {
	var userModel model.UserModel

	err := r.db.First(&userModel, "email = ?", email).Error

	if err != nil {
		// Se o erro for "record not found", significa que o email não existe,
		// então retornamos false sem erro
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *GormUserRepository) Update(user entity.User) error {
	return r.
		db.
		Model(&model.UserModel{}).
		Where("id = ?", user.ID).
		Updates(mapper.ToUserModel(user)).
		Error
}

func (r *GormUserRepository) Delete(id string) error {
	return r.
		db.
		Delete(&model.UserModel{}, "id = ?", id).
		Error
}

func (r *GormUserRepository) Restore(id string) error {
	return r.
		db.
		Unscoped().
		Model(&model.UserModel{}).
		Where("id = ?", id).
		Update("deleted_at", nil).
		Error
}
