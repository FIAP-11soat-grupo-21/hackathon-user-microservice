package cognito

import (
	"context"
	"errors"
	"log"
	"user_microservice/internal/common/config/env"
	"user_microservice/internal/core/dto"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AWSCognitoAuthService struct {
	Client     *cognitoidentityprovider.Client
	UserPoolId string
}

func NewAWSCognitoAuthService() *AWSCognitoAuthService {
	cfg := env.GetConfig()

	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.AWS.Region),
	)

	if err != nil {
		panic(err)
	}

	client := cognitoidentityprovider.NewFromConfig(awsConfig)

	return &AWSCognitoAuthService{
		Client:     client,
		UserPoolId: cfg.AWS.Cognito.UserPoolId,
	}
}

func (c *AWSCognitoAuthService) RegisterUser(user dto.RegisterUserDTO) error {
	ctx := context.Background()

	cognitoCustomer, err := c.Client.AdminCreateUser(ctx, &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId: aws.String(c.UserPoolId),
		Username:   aws.String(user.Email),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	})

	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return err
	}

	_, err = c.Client.AdminSetUserPassword(ctx, &cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: aws.String(c.UserPoolId),
		Username:   aws.String(user.Email),
		Password:   aws.String(user.Password),
		Permanent:  true,
	})

	if err != nil {
		log.Printf("Failed to set user password: %v", err)
		return err
	}

	if cognitoCustomer == nil || cognitoCustomer.User == nil {
		log.Printf("AdminCreateUser returned unexpected result: %#v", cognitoCustomer)
		return errors.New("failed to retrieve cognito user information")
	}

	cognitoSubId := ""
	for _, attr := range cognitoCustomer.User.Attributes {
		if attr.Name != nil && *attr.Name == "sub" && attr.Value != nil {
			cognitoSubId = *attr.Value
			break
		}
	}

	if cognitoSubId == "" {
		log.Printf("Failed to retrieve cognito sub ID for user: %s", user.Email)
		return errors.New("cognito sub id not found")
	}

	return nil
}

func (c *AWSCognitoAuthService) UpdateUserEmail(oldEmail, newEmail string) error {
	ctx := context.Background()

	_, err := c.Client.AdminUpdateUserAttributes(ctx, &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserPoolId: aws.String(c.UserPoolId),
		Username:   aws.String(oldEmail),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(newEmail),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
	})

	if err != nil {
		log.Printf("Failed to update user email: %v", err)
		return err
	}

	return nil
}

func (c *AWSCognitoAuthService) UpdateUserPassword(email, newPassword string) error {
	ctx := context.Background()
	_, err := c.Client.AdminSetUserPassword(ctx, &cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: aws.String(c.UserPoolId),
		Username:   aws.String(email),
		Password:   aws.String(newPassword),
		Permanent:  true,
	})

	if err != nil {
		log.Printf("Failed to update user password: %v", err)
		return err
	}

	return nil
}

func (c *AWSCognitoAuthService) DeleteUser(email string) error {
	ctx := context.Background()

	_, err := c.Client.AdminDisableUser(ctx, &cognitoidentityprovider.AdminDisableUserInput{
		UserPoolId: &c.UserPoolId,
		Username:   &email,
	})

	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}

	return nil
}

func (c *AWSCognitoAuthService) RestoreUser(email string) error {
	ctx := context.Background()

	_, err := c.Client.AdminEnableUser(ctx, &cognitoidentityprovider.AdminEnableUserInput{
		UserPoolId: &c.UserPoolId,
		Username:   &email,
	})

	if err != nil {
		log.Printf("Failed to restore user: %v", err)
		return err
	}

	return nil
}
