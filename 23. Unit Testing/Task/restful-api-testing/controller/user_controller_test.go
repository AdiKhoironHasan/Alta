package controller

import (
	"encoding/json"
	"logging-and-jwt/config"
	"logging-and-jwt/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/tj/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		assert.NoError(t, GetUsersController(c))
		assert.Equal(t, testCase.expectCode, rec.Code)
		body := rec.Body.String()

		var user UserResponse
		err := json.Unmarshal([]byte(body), &user)

		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, testCase.sizeData, len(user.Data))
	}
}
