package schema

type CreateUserSchema struct {
	Name     string `json:"name" example:"João Silva" binding:"required"`
	Email    string `json:"email" example:"joao@example.com" binding:"required"`
	Password string `json:"password" example:"password" binding:"required"`
}

type UpdateUserSchema struct {
	Name  string `json:"name" example:"João Silva" binding:"required"`
	Email string `json:"email" example:"joao@example.com" binding:"required"`
}

type UpdateUserPasswordSchema struct {
	NewPassword string `json:"new_password" example:"newpassword" binding:"required"`
}

type UserResponseSchema struct {
	ID    string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name  string `json:"name" example:"João Silva"`
	Email string `json:"email" example:"joao@example.com"`
}

type UserNotFoundErrorSchema struct {
	Error string `json:"error" example:"User not found"`
}

type UserAlreadyExistsErrorSchema struct {
	Error string `json:"error" example:"User already exists"`
}

type InvalidUserDataErrorSchema struct {
	Error string `json:"error" example:"Invalid user data"`
}

type ErrorMessageSchema struct {
	Error string `json:"error" example:"Internal server error"`
}
