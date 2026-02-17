package value_object

import (
	"net/mail"
	"strings"
	"user_microservice/internal/core/domain/exception"
)

type Email struct {
	value string
}

func (e *Email) Value() string {
	return e.value
}

func NewEmail(email string) (Email, error) {
	if !isValidEmail(email) {
		return Email{}, &exception.InvalidUserDataException{
			Message: "Invalid email address: " + email,
		}
	}

	return Email{value: strings.ToLower(email)}, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
