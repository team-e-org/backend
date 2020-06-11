package infrastructure

import (
	"app/models/db"
	"app/repository"
	"database/sql"
)

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(user *db.User) error {
	return nil
}

func (u *User) GetUser(userID int) (*db.User, error) {
	return nil, nil
}
