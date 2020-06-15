package infrastructure

import (
	"app/models"
	"app/repository"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (u *User) CreateUser(name string, email string, icon string, password string) (*models.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	const query = `
INSERT INTO users (email, password, name, icon, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`
	now := time.Now()

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(email, hashedPassword, name, icon, now, now)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:             int(id),
		Email:          email,
		HashedPassword: hashedPassword,
		Icon:           icon,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	return user, nil
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
