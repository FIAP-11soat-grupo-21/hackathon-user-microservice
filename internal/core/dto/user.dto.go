package dto

type CreateUserDTO struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserDTO struct {
	ID    string
	Name  string
	Email string
}

type UpdateUserPasswordDTO struct {
	ID          string
	NewPassword string
}
