package exception

type UserNotFoundException struct {
	Message string
}

func (err *UserNotFoundException) Error() string {
	if err == nil {
		return "User not found"
	}
	if err.Message != "" {
		return err.Message
	}
	return "User not found"
}

type UserAlreadyExistsException struct {
	Message string
}

func (e *UserAlreadyExistsException) Error() string {
	if e.Message == "" {
		return "User already exists"
	}
	return e.Message
}

type InvalidUserDataException struct {
	Message string
}

func (e *InvalidUserDataException) Error() string {
	if e.Message == "" {
		return "Invalid user data"
	}

	return e.Message
}
