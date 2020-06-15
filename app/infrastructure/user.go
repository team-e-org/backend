package infrastructure

import (
	"app/models"
	"app/repository"
	"database/sql"
	"errors"
)

var ErrEmailNotFound = errors.New("email doesn't exist in the database")

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(user *models.User) error {
	return nil
}

func (u *User) UpdateUser(user *models.User) error {
	return nil
}

func (u *User) DeleteUser(userID int) error {
	return nil
}

func (u *User) GetUser(userID int) (*models.User, error) {
	return nil, nil
}

func (u *User) GetUserByEmail(email string) (*models.User, error) {
	const query = `
SELECT u.id, u.email, u.password, u.icon, u.created_at, u.updated_at
FROM users u
WHERE u.email = ?;
`
	row := u.DB.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.HashedPassword,
		&user.Icon,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, ErrEmailNotFound
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
