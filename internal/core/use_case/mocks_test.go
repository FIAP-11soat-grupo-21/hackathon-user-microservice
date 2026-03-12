package use_case

import (
	"user_microservice/internal/core/domain/entity"
	"user_microservice/internal/core/dto"
)

type fakeUserRepository struct {
	existsByEmailResult bool
	existsByEmailErr    error
	insertErr           error
	insertCalled        bool

	findByIDResult entity.User
	findByIDErr    error

	listAllResult []entity.User
	listAllErr    error

	updateErr    error
	updateCalled bool

	deleteErr    error
	deleteCalled bool

	restoreErr    error
	restoreCalled bool
}

func (f *fakeUserRepository) ExistsByEmail(_ string) (bool, error) {
	return f.existsByEmailResult, f.existsByEmailErr
}

func (f *fakeUserRepository) Insert(_ entity.User) error {
	f.insertCalled = true
	return f.insertErr
}

func (f *fakeUserRepository) ListAll() ([]entity.User, error) {
	return f.listAllResult, f.listAllErr
}

func (f *fakeUserRepository) FindByID(_ string) (entity.User, error) {
	return f.findByIDResult, f.findByIDErr
}

func (f *fakeUserRepository) FindByEmail(_ string) (entity.User, error) { return entity.User{}, nil }

func (f *fakeUserRepository) Update(_ entity.User) error {
	f.updateCalled = true
	return f.updateErr
}

func (f *fakeUserRepository) Delete(_ string) error {
	f.deleteCalled = true
	return f.deleteErr
}

func (f *fakeUserRepository) Restore(_ string) error {
	f.restoreCalled = true
	return f.restoreErr
}

type fakeAuthService struct {
	hashPasswordResult string
	hashPasswordErr    error

	registerUserCalled bool
	registerUserErr    error

	updateEmailCalled bool
	updateEmailErr    error

	updatePasswordCalled bool
	updatePasswordErr    error

	restoreUserCalled bool
	restoreUserErr    error

	deleteUserCalled bool
	deleteUserErr    error
}

func (f *fakeAuthService) HashPassword(_ string) (string, error) {
	return f.hashPasswordResult, f.hashPasswordErr
}

func (f *fakeAuthService) RegisterUser(_ dto.RegisterUserDTO) error {
	f.registerUserCalled = true
	return f.registerUserErr
}

func (f *fakeAuthService) UpdateUserEmail(_, _ string) error {
	f.updateEmailCalled = true
	return f.updateEmailErr
}

func (f *fakeAuthService) UpdateUserPassword(_, _ string) error {
	f.updatePasswordCalled = true
	return f.updatePasswordErr
}

func (f *fakeAuthService) RestoreUser(_ string) error {
	f.restoreUserCalled = true
	return f.restoreUserErr
}

func (f *fakeAuthService) DeleteUser(_ string) error {
	f.deleteUserCalled = true
	return f.deleteUserErr
}
