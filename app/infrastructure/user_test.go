package infrastructure

import (
	"app/models"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var sqlDB *sql.DB
var mock sqlmock.Sqlmock
var now time.Time

func mockDBHandlingWrapper(m *testing.M) int {
	var err error
	sqlDB, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	now = time.Now()

	defer sqlDB.Close()

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(mockDBHandlingWrapper(m))
}

func TestInsertUser(t *testing.T) {
	user := &models.User{
		ID:             0,
		Name:           "test user",
		Email:          "test@test.com",
		Icon:           "test icon",
		HashedPassword: "password",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO users (name, email, password, icon) VALUES (?, ?, ?, ?)"))
	prepare.ExpectExec().
		WithArgs(user.Name, user.Email, user.HashedPassword, user.Icon).
		WillReturnResult(sqlmock.NewResult(0, 1))

	users := NewUserRepository(sqlDB)
	err := users.CreateUser(user)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestInsertUserError(t *testing.T) {
	user := &models.User{
		ID:             0,
		Name:           "test user",
		Email:          "test@test.com",
		Icon:           "test icon",
		HashedPassword: "password",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO users (name, email, password, icon) VALUES (?, ?, ?, ?)"))
	prepare.ExpectExec().
		WithArgs(user.Name, user.Email, user.HashedPassword, user.Icon).
		WillReturnError(fmt.Errorf("some error"))

	users := NewUserRepository(sqlDB)
	err := users.CreateUser(user)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestUpdateUser(t *testing.T) {
	user := &models.User{
		ID:             0,
		Name:           "test user",
		Email:          "test@test.com",
		Icon:           "test icon",
		HashedPassword: "password",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, password = ?, icon = ? WHERE id = ?"))
	prepare.ExpectExec().
		WithArgs(user.Name, user.Email, user.HashedPassword, user.Icon, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	users := NewUserRepository(sqlDB)
	err := users.UpdateUser(user)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestUpdateUserError(t *testing.T) {
	user := &models.User{
		ID:             0,
		Name:           "test user",
		Email:          "test@test.com",
		Icon:           "test icon",
		HashedPassword: "password",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("UPDATE users SET name = ?, email = ?, password = ?, icon = ? WHERE id = ?"))
	prepare.ExpectExec().
		WithArgs(user.Name, user.Email, user.HashedPassword, user.Icon, user.ID).
		WillReturnError(fmt.Errorf("some error"))

	users := NewUserRepository(sqlDB)
	err := users.UpdateUser(user)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestDeleteUser(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM users WHERE id = ?"))

	prepare.ExpectExec().
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	users := NewUserRepository(sqlDB)
	err := users.DeleteUser(userID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestDeleteUserError(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM users WHERE id = ?"))

	prepare.ExpectExec().
		WithArgs(userID).
		WillReturnError(fmt.Errorf("some error"))

	users := NewUserRepository(sqlDB)
	err := users.DeleteUser(userID)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetUser(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare("SELECT u.id, u.Name, u.email, u.password, u.icon. u.created_at, u.updated_at FROM users u WHERE u.id = ?")
	prepare.ExpectQuery().
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "icon", "created_at", "updated_at"}).
			AddRow(0, "test user", "test@test.com", "password", "test icon", now, now))

	users := NewUserRepository(sqlDB)
	_, err := users.GetUser(userID)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetUserError(t *testing.T) {
	userID := 0

	prepare := mock.ExpectPrepare("SELECT u.id, u.Name, u.email, u.password, u.icon. u.created_at, u.updated_at FROM users u WHERE u.id = ?")
	prepare.ExpectQuery().
		WithArgs(userID).
		WillReturnError(fmt.Errorf("some error"))

	users := NewUserRepository(sqlDB)
	_, err := users.GetUser(0)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	email := "test@test.com"

	prepare := mock.ExpectPrepare("SELECT u.id, u.name, u.email, u.password, u.icon. u.created_at, u.updated_at FROM users u WHERE u.email = ?")
	prepare.ExpectQuery().
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "icon", "created_at", "updated_at"}).
			AddRow(0, "test user", "test@test.com", "password", "test icon", now, now))

	users := NewUserRepository(sqlDB)
	_, err := users.GetUserByEmail(email)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}

func TestGetUserByEmailError(t *testing.T) {
	email := "test@test.com"

	prepare := mock.ExpectPrepare("SELECT u.id, u.name, u.email, u.password, u.icon. u.created_at, u.updated_at FROM users u WHERE u.email = ?")
	prepare.ExpectQuery().
		WithArgs(email).
		WillReturnError(fmt.Errorf("some error"))

	users := NewUserRepository(sqlDB)
	_, err := users.GetUserByEmail(email)
	if err == nil {
		t.Fatalf("An error occurred: %v\n", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Unfulfilled expectations error: %v\n", err)
	}
}
