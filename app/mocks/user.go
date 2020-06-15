package mocks

import (
	"app/models"
	"app/repository"
	"errors"
)

type UserMock struct {
	ExpectedUser *models.User
}

func NewUserRepository() repository.UserRepository {
	return &UserMock{}
}

func (m *UserMock) CreateUser(user *models.User) error {
	m.ExpectedUser = user
	return nil
}

func (m *UserMock) UpdateUser(user *models.User) error {
	if m.ExpectedUser == nil {
		return noUserError()
	}
	m.ExpectedUser = user
	return nil
}

func (m *UserMock) DeleteUser(userID int) error {
	if m.ExpectedUser == nil {
		return noUserError()
	}
	m.ExpectedUser = nil
	return nil
}

func (m *UserMock) GetUser(userID int) (*models.User, error) {
	if m.ExpectedUser.ID != userID {
		return nil, noUserError()
	}
	return m.ExpectedUser, nil
}

func (m *UserMock) GetUserByEmail(email string) (*models.User, error) {
	if m.ExpectedUser.Email != email {
		return nil, noUserError()
	}
	return m.ExpectedUser, nil
}

func noUserError() error {
	return errors.New("An error occurred, the user does not exist")
}
