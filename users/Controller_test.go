package user

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockDB   = map[string]*User{}
	userJSON = `{"UserName":"Jon Snow", "Password": "test"}`
)

func TestNewController(t *testing.T) {

}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(req, recorder)
	ctrl := &Controller{service: Service{repository: Repository{data: map[string]User{}}}}

	if assert.NoError(t, ctrl.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(t, userJSON, recorder.Body.String())
	}
}

func TestLogin(t *testing.T) {

}
