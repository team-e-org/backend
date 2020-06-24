package server

import (
	"app/authz"
	"app/db"
	"app/goldenfiles"
	"app/models"
	"app/ptr"
	helpers "app/testutils"
	"app/testutils/dbdata"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUserBoards(t *testing.T) {
	var cases = []struct {
		Desc   string
		Code   int
		userID int
		boards []*models.Board
	}{
		{
			"single board",
			200,
			1,
			[]*models.Board{board1()},
		},
		{
			"two board",
			200,
			1,
			[]*models.Board{board1(), board2()},
		},
		{
			"two boards with one private",
			200,
			1,
			[]*models.Board{board1(), board2(), privateBoard()},
		},
		{
			"not found user",
			404,
			2,
			[]*models.Board{},
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			for _, b := range c.boards {
				data.Boards().CreateBoard(b)
			}

			attachHandlers(router, data, authz.NewAuthLayerMock(data))
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d/boards", c.userID), nil)

			router.ServeHTTP(recorder, req)
			body := recorder.Body.Bytes()

			assert.Equal(t, c.Code, recorder.Code, "Status code should match reference")
			expected := goldenfiles.UpdateAndOrRead(t, body)
			assert.Equal(t, expected, body, "Response body should match golden file")
		})
	}
}

func TestSignUp(t *testing.T) {
	var cases = []struct {
		Desc        string
		Code        int
		requestBody string
	}{
		{
			"success",
			201,
			`{"name": "test user","email": "user@example.com","password": "pa$$wOrd12345"}`,
		},
		{
			"invalid email",
			400,
			`{"name": "test user","email": "userexample.com","password": "pa$$wOrd12345"}`,
		},
		{
			"short password",
			400,
			`{"name": "test user","email": "user@example.com","password": "p$AO15"}`,
		},
		{
			"only lower password",
			400,
			`{"name": "test user","email": "user@example.com","password": "aaaaaaaaaaa"}`,
		},
		{
			"only capital password",
			400,
			`{"name": "test user","email": "user@example.com","password": "AAAAAAAAAAA"}`,
		},
		{
			"only number password",
			400,
			`{"name": "test user","email": "user@example.com","password": "12345678910"}`,
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			attachHandlers(router, data, authz.NewAuthLayerMock(data))
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(
				http.MethodPost,
				"/users/sign-up",
				ioutil.NopCloser(strings.NewReader(c.requestBody)))

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.Code, recorder.Code, "Status code should match reference")
			// cannot assert response body because token changes every request
		})
	}
}

func TestSignIn(t *testing.T) {
	var cases = []struct {
		Desc        string
		Code        int
		requestBody string
	}{
		{
			"success",
			200,
			`{"email": "current_user@email.com","password": "password"}`,
		},
		{
			"wrong email",
			401,
			`{"email": "wrong@email.com","password": "password"}`,
		},
		{
			"wrong password",
			401,
			`{"email": "sss@email.com","password": "worngPasword"}`,
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			data.Users().CreateUser(dbdata.BaseUser)

			al := authz.NewAuthLayerMock(data)

			attachHandlers(router, data, al)
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(
				http.MethodPost,
				"/users/sign-in",
				ioutil.NopCloser(strings.NewReader(c.requestBody)))

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.Code, recorder.Code, "Status code should match reference")
			// cannot assert response body because token changes every request
		})
	}
}

func board1() *models.Board {
	return &models.Board{
		ID:          1,
		UserID:      1,
		Name:        "test name1",
		Description: ptr.NewString("test description1"),
	}
}

func board2() *models.Board {
	return &models.Board{
		ID:          2,
		UserID:      1,
		Name:        "test name2",
		Description: ptr.NewString("test description2"),
	}
}

func privateBoard() *models.Board {
	return &models.Board{
		ID:          3,
		UserID:      1,
		Name:        "test name private",
		Description: ptr.NewString("test description private"),
		IsPrivate:   true,
	}
}
