package encrypt

import (
	"user_microservice/internal/common/config/env"

	"golang.org/x/crypto/bcrypt"
)

func ByBcrypt(password string) (string, error) {
	cfg := env.GetConfig()
	cost := cfg.PasswordSalt

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
