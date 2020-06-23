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
	authToken := helpers.BasicUserToken
	tokenData := &authz.TokenData{
		UserData: dbdata.BaseUser,
	}

	var cases = []struct {
		Desc        string
		Code        int
		requestBody string
	}{
		{
			"success",
			201,
			`{"name": "new board"}`,
		},
		{
			"success with more params",
			201,
			`{"name": "new board", "description": "test description", "isPrivate": true}`,
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			al := &authz.AuthLayerMock{}
			al.ExpectedTokenData = tokenData

			attachReqAuth(router, data, al)
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/boards", ioutil.NopCloser(strings.NewReader(c.requestBody)))
			helpers.SetAuthTokenHeader(req, authToken)

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
