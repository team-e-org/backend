package mocks

import (
	"app/models"
	"app/repository"
)

type UserMock struct {
	ExpectedUser *models.User
}

func NewUserRepository() repository.UserRepository {
	return &UserMock{}
}

func (m *UserMock) AddUser(user *db.User) error {
	m.ExpectedUser = user
	return nil
}

func (m *UserMock) GetUser(userID int) (*models.User, error) {
	return m.ExpectedUser, nil
}
