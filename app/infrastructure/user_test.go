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

func TestUserRepositorySQLMock(t *testing.T) {
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
