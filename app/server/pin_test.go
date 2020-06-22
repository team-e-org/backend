package server

import (
	"app/authz"
	"app/db"
	"app/goldenfiles"
	"app/mocks"
	"app/models"
	"app/ptr"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	helpers "app/testutils"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestServePin(t *testing.T) {
	var cases = []struct {
		Desc  string
		Code  int
		pinID int
		pin   *models.Pin
	}{
		{
			"single pin",
			200,
			1,
			pin1(),
		},
		{
			"no pin",
			404,
			2,
			pin1(),
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			mockPinRepository := mocks.NewPinRepository()
			mockPinRepository.CreatePin(c.pin, 0)
			data.Pins = mockPinRepository

			attachHandlers(router, data, authz.NewAuthLayer(data))
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/pins/%d", c.pinID), nil)

			router.ServeHTTP(recorder, req)
			body := recorder.Body.Bytes()

			assert.Equal(t, c.Code, recorder.Code, "Status code should match reference")
			expected := goldenfiles.UpdateAndOrRead(t, body)
			assert.Equal(t, expected, body, "Response body should match golden file")
		})
	}
}

func pin1() *models.Pin {
	return &models.Pin{
		ID:          1,
		UserID:      ptr.NewInt(0),
		Title:       "test title",
		Description: ptr.NewString("test description"),
		URL:         ptr.NewString("test url"),
		ImageURL:    "test image url",
		IsPrivate:   false,
	}
}
