package infrastructure

import (
	"app/models"
	"database/sql"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

var sqlDB *sql.DB
var mock sqlmock.Sqlmock

func mockDBHandlingWrapper(m *testing.M) int {
	var err error
	sqlDB, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(mockDBHandlingWrapper(m))
}

func TestInsertUser(t *testing.T) {
	now := time.Now()
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
}

func TestGetUser(t *testing.T) {
	now := time.Now()
	prepare := mock.ExpectPrepare("SELECT u.id, u.Name, u.email, u.password, u.icon. u.created_at, u.updated_at FROM users u WHERE u.id = ?")
	prepare.ExpectQuery().WithArgs(0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "icon", "created_at", "updated_at"}).
			AddRow(0, "test user", "test@test.com", "password", "test icon", now, now))

	users := NewUserRepository(sqlDB)
	_, err := users.GetUser(0)
	if err != nil {
		t.Fatalf("An error occurred: %v\n", err)
	}
}
