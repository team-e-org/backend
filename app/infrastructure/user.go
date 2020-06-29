package infrastructure

import (
	"app/helpers"
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

func (u *User) CreateUser(user *models.User) (*models.User, error) {
	const query = `
    INSERT INTO users (name, email, password, icon) VALUES (?, ?, ?, ?);
    `

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(user.Name, user.Email, user.HashedPassword, user.Icon)
	err = helpers.CheckDBExecError(result, err)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}

func (u *User) UpdateUser(user *models.User) error {
	const query = `
    UPDATE users SET name = ?, email = ?, password = ?, icon = ? WHERE id = ?;
    `

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Name, user.Email, user.HashedPassword, user.Icon, user.ID)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser(userID int) error {
	const query = `
    DELETE FROM users WHERE id = ?;
    `

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(userID)
	if err = helpers.CheckDBExecError(result, err); err != nil {
		return err
	}

	return nil
}

func (u *User) GetUser(userID int) (*models.User, error) {
	const query = `
    SELECT u.id, u.Name, u.email, u.password, u.icon, u.created_at, u.updated_at FROM users u WHERE u.id = ?;
    `

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userID)

	user := &models.User{}
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Icon,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetUserByEmail(email string) (*models.User, error) {
	const query = `
SELECT u.id, u.name, u.email, u.password, u.icon, u.created_at, u.updated_at
FROM users u
WHERE u.email = ?;
`

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(email)

	user := &models.User{}
	err = row.Scan(
		&user.ID,
		&user.Name,
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
