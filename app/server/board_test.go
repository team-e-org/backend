package server

import (
	"app/authz"
	"app/db"
	"app/goldenfiles"
	"app/mocks"
	"app/models"
	helpers "app/testutils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateBoard(t *testing.T) {
	var cases = []struct {
		Desc        string
		Code        int
		requestBody string
		currentUser *models.User
	}{
		{
			"success",
			201,
			`{"name": "new board"}`,
			currentUser(),
		},
		{
			"success with more params",
			201,
			`{"name": "new board", "description": "test description", "isPrivate": true}`,
			currentUser(),
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			mockUserRepository := mocks.NewUserRepository()
			mockUserRepository.CreateUser(c.currentUser)
			data.Users = mockUserRepository

			al := authz.NewAuthLayer(data)
			token, _ := al.AuthenticateUser(c.currentUser.Email, "password")

			attachReqAuth(router, data, al)
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/boards", ioutil.NopCloser(strings.NewReader(c.requestBody)))
			if token != "" {
				helpers.SetAuthTokenHeader(req, token)
			}

			router.ServeHTTP(recorder, req)
			body := recorder.Body.Bytes()

			assert.Equal(t, c.Code, recorder.Code, "Status code should match reference")
			expected := goldenfiles.UpdateAndOrRead(t, body)
			assert.Equal(t, expected, body, "Response body should match golden file")
		})
	}
}

func currentUser() *models.User {
	return &models.User{
		ID:             1,
		Name:           "current user",
		Email:          "current_user@email.com",
		Icon:           "test icon",
		HashedPassword: "$2a$10$SOWUFP.hkVI0CrCJyfh5vuf/Gu.SDpv6Y2DYZ/Dbwyr.AKtlAldFe",
	}
}
