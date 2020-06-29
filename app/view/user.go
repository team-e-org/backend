package view

import (
	"app/helpers"
	"app/models"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	Icon     string `json:"icon"`
}

func NewUser(user *models.User) *User {
	u := &User{
		user.ID,
		user.Name,
		user.Email,
		user.HashedPassword,
		user.Icon,
	}

	return u
}

func NewUserModel(user *User) (*models.User, error) {
	hashedPassword := ""
	if user.Password != "" {
		p, err := helpers.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		hashedPassword = p
	}

	u := &models.User{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Icon:           user.Icon,
		HashedPassword: hashedPassword,
	}

	return u, nil
}
