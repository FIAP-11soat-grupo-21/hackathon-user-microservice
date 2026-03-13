package factory

import (
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	testFactoryFunction(t,
		"Should return non-nil user repository",
		func() interface{} { return NewUserRepository() },
		"repository")
}
