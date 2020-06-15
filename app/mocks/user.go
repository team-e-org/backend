package mocks

import (
	"app/models"
	"app/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserMock struct {
	ExpectedUser *models.User
}

func NewUserRepository() repository.UserRepository {
	return &UserMock{}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (m *UserMock) CreateUser(name string, email string, icon string, password string) (*models.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user := &models.User{
		ID:             0,
		Email:          email,
		HashedPassword: hashedPassword,
		Icon:           icon,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	m.ExpectedUser = user

	return m.ExpectedUser, nil
}

func (m *UserMock) GetUser(userID int) (*models.User, error) {
	return m.ExpectedUser, nil
}

func (m *UserMock) GetUserByEmail(email string) (*models.User, error) {
	return m.ExpectedUser, nil
}
