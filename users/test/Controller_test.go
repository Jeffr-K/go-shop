package test

import (
	"github.com/labstack/echo/v4"
	"github.com/lambsroad/go-shop/users"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userJSON = `{"UserName":"Jon Snow", "Password": "test"}`
)

func TestCreateUser(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)

	service := user.Service{}
	ctrl := user.NewController(service)

	if assert.NoError(t, ctrl.CreateUser(context)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, userJSON, recorder.Body.String())
	}
}

func TestLogin(t *testing.T) {

}
