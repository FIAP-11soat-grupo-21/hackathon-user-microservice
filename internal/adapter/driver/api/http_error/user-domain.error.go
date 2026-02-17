package http_error

import (
	"net/http"
	"user_microservice/internal/core/domain/exception"

	"github.com/gin-gonic/gin"
)

func HandleDomainErrors(err error, ctx *gin.Context) bool {
	var status int

	switch err.(type) {
	case *exception.UserNotFoundException:
		status = http.StatusNotFound
	case *exception.UserAlreadyExistsException:
		status = http.StatusConflict
	case *exception.InvalidUserDataException:
		status = http.StatusBadRequest
	}

	if status != 0 {
		ctx.JSON(status, gin.H{"error": err.Error()})
		return true
	}

	return false
}
