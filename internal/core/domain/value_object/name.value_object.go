package value_object

import (
	"strings"
	"user_microservice/internal/core/domain/exception"
)

type Name struct {
	value string
}

func (n *Name) Value() string {
	return n.value
}

func NewName(name string) (Name, error) {
	name = strings.TrimSpace(name)

	if len(name) < 3 {
		return Name{}, &exception.InvalidUserDataException{
			Message: "name must have at least 3 characters",
		}
	}

	if len(name) > 100 {
		return Name{}, &exception.InvalidUserDataException{
			Message: "name must have at most 100 characters",
		}
	}

	return Name{value: name}, nil
}
