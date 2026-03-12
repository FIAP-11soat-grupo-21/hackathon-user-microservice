package factory

import (
	"testing"
)

func TestNewAuthService(t *testing.T) {
	testFactoryFunction(t,
		"Should return non-nil auth service",
		func() interface{} { return NewAuthService() },
		"auth service")
}
