package model

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        string          `gorm:"column:id;primaryKey;size:36"`
	Name      string          `gorm:"column:name;not null;size:255"`
	Email     string          `gorm:"column:email;not null;size:255;unique"`
	Password  string          `gorm:"column:password;not null;size:255"`
	CreatedAt time.Time       `gorm:"column:created_at;not null;type:timestamp(6)"`
	UpdatedAt time.Time       `gorm:"column:updated_at;not null;type:timestamp(6)"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp(6)"`
}

func (UserModel) TableName() string { return "users" }
