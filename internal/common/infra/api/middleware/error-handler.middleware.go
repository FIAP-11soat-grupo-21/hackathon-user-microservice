package middleware

import (
	"net/http"
	"user_microservice/internal/adapter/driver/api/http_error"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			errorHasBinHandled := http_error.HandleDomainErrors(err, ctx)

			if !errorHasBinHandled {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Internal server error",
					"details": err.Error(),
				})
			}

			ctx.Abort()
		}
	}
}
