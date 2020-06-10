package repository

import "app/models/db"

type UserRepository interface {
	AddUser(user *db.User) error
	GetUser(userID int) (*db.User, error)
}
