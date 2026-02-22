package factory

import (
	"user_microservice/internal/adapter/driven/aws/cognito"
	"user_microservice/internal/core/domain/port"
)

func NewAuthService() port.IAuthService {
	return cognito.NewAWSCognitoAuthService()
}
