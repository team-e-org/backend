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
			500, // testでは500になる。sql.ErrNoRowsのため
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

			attachHandlers(router, data, authz.NewAuthLayerMock(data))
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

func TestServePinsInBoard(t *testing.T) {
	var cases = []struct {
		Desc    string
		Code    int
		boardID int
		pins    []*models.Pin
		page    int
	}{
		{
			"single pin",
			200,
			1,
			[]*models.Pin{pin1()},
			1,
		},
		{
			"two pin",
			200,
			1,
			[]*models.Pin{pin1(), pin2()},
			1,
		},
		{
			"private pin",
			200,
			1,
			[]*models.Pin{pin1(), pin2(), privatePin()},
			1,
		},
		{
			"no pin",
			200,
			1,
			[]*models.Pin{privatePin()},
			1,
		},
	}

	for _, c := range cases {
		t.Run(helpers.TableTestName(c.Desc), func(t *testing.T) {
			router := mux.NewRouter()
			data := db.NewRepositoryMock()

			mockPinRepository := mocks.NewPinRepository()
			for _, p := range c.pins {
				mockPinRepository.CreatePin(p, c.boardID)
			}
			data.Pins = mockPinRepository

			attachHandlers(router, data, authz.NewAuthLayerMock(data))
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/boards/%d/pins?page=%d", c.boardID, c.page), nil)

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

func pin2() *models.Pin {
	return &models.Pin{
		ID:          2,
		UserID:      ptr.NewInt(0),
		Title:       "test title 2",
		Description: ptr.NewString("test description2"),
		URL:         ptr.NewString("test url2"),
		ImageURL:    "test image url2",
		IsPrivate:   false,
	}
}

func privatePin() *models.Pin {
	return &models.Pin{
		ID:          3,
		UserID:      ptr.NewInt(1),
		Title:       "test title  private",
		Description: ptr.NewString("test description private"),
		URL:         ptr.NewString("test url private"),
		ImageURL:    "test image url private",
		IsPrivate:   true,
	}
}
