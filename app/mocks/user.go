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

func (m *UserMock) GetUser(userID int) (*models.User, error) {
	return m.ExpectedUser, nil
}

func (m *UserMock) GetUserByEmail(email string) (*models.User, error) {
	if m.ExpectedUser.Email != email {
		return nil, errors.New("user not found")
	}
	return m.ExpectedUser, nil
}
