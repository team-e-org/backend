package server

import (
	"app/authz"
	"app/db"
	"app/goldenfiles"
	"app/models"
	helpers "app/testutils"
	"app/testutils/dbdata"
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
		Desc          string
		Code          int
		requestBody   string
		currentUser   *models.User
		loginPassword string
	}{
		{
			"success",
			201,
			`{"userId": 1, "name": "new board"}`,
			dbdata.BaseUser,
			dbdata.BaseUserPassword,
		},
		{
			"success with more params",
			201,
			`{"userId": 1, "name": "new board", "description": "test description", "isPrivate": true}`,
			dbdata.BaseUser,
			dbdata.BaseUserPassword,
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			data.Users().CreateUser(c.currentUser)

			al := authz.NewAuthLayerMock(data)
			token, _ := al.AuthenticateUser(c.currentUser.Email, c.loginPassword)

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
