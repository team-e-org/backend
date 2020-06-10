package mocks

import (
	"app/models/db"
	"app/repository"
)

type UserMock struct {
	ExpectedUser *db.User
}

func NewUserRepository() repository.UserRepository {
	return &UserMock{}
}

func (m *UserMock) AddUser(user *db.User) error {
	m.ExpectedUser = user
	return nil
}

func (m *UserMock) GetUser(userID int) (*db.User, error) {
	return m.ExpectedUser, nil
}
