package value_object

import (
	"regexp"
	"user_microservice/internal/common/pkg/encrypt"
	"user_microservice/internal/core/domain/exception"
)

var (
	lowercase   = regexp.MustCompile(`[a-z]`)
	uppercase   = regexp.MustCompile(`[A-Z]`)
	digit       = regexp.MustCompile(`[0-9]`)
	specialChar = regexp.MustCompile(`[^A-Za-z0-9]`)
)

type Password struct {
	value string
}

func (p *Password) Value() string {
	return p.value
}

func NewPassword(password string) (Password, error) {
	hasLower := lowercase.MatchString(password)
	hasUpper := uppercase.MatchString(password)
	hasDigit := digit.MatchString(password)
	hasSpecial := specialChar.MatchString(password)

	if len(password) < 8 ||
		!hasLower ||
		!hasUpper ||
		!hasDigit ||
		!hasSpecial {
		return Password{}, &exception.InvalidUserDataException{
			Message: "Password must be at least 8 characters long and include uppercase, lowercase, numbers, and special characters",
		}
	}

	encryptedPassword, err := encrypt.ByBcrypt(password)

	if err != nil {
		return Password{}, err
	}

	return Password{value: encryptedPassword}, nil
}
