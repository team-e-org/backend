package view

import "app/models"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Icon  string `json:"icon"`
}

func NewUser(user *models.User) *User {
	u := &User{
		user.ID,
		user.Name,
		user.Email,
		user.Icon,
	}

	return u
}
