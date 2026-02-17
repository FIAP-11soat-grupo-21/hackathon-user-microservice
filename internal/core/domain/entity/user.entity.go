package entity

import (
	"user_microservice/internal/common/pkg/identity"
	"user_microservice/internal/core/domain/exception"
	"user_microservice/internal/core/domain/value_object"
)

type User struct {
	ID       string
	Name     value_object.Name
	Email    value_object.Email
	Password value_object.Password
}

func NewUser(id, name, email, password string) (*User, error) {
	newName, err := value_object.NewName(name)

	if err != nil {
		return nil, err
	}

	newEmail, err := value_object.NewEmail(email)

	if err != nil {
		return nil, err
	}

	newPassword, err := value_object.NewPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Name:     newName,
		Email:    newEmail,
		Password: newPassword,
	}, nil
}

func NewUserWithoutPassword(id, name, email string) (*User, error) {
	newName, err := value_object.NewName(name)

	if err != nil {
		return nil, err
	}

	newEmail, err := value_object.NewEmail(email)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:    id,
		Name:  newName,
		Email: newEmail,
	}, nil
}

func (u *User) SetName(name string) error {
	newName, err := value_object.NewName(name)
	if err != nil {
		return err
	}

	u.Name = newName
	return nil
}

func (u *User) SetEmail(email string) error {
	newEmail, err := value_object.NewEmail(email)
	if err != nil {
		return err
	}

	u.Email = newEmail
	return nil
}

func (u *User) SetPassword(password string) error {
	newPassword, err := value_object.NewPassword(password)
	if err != nil {
		return err
	}

	u.Password = newPassword
	return nil
}

func ValidateID(id string) error {
	if !identity.IsValidUUID(id) {
		return &exception.InvalidUserDataException{
			Message: "Invalid user ID",
		}
	}

	return nil
}

func (u *User) IsEmpty() bool {
	return u.ID == ""
}
