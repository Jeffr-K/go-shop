package go_shop

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	user "github.com/lambsroad/go-shop/users"
	"log"
	"net/http"
)

const _defaultPort = 4000

type Server struct {
	userController user.Controller
}

func NewDefaultServer() *Server {
	data := map[string]user.User{}
	userRepository := user.NewRepository(data)
	userService := user.NewService(*userRepository)
	userController := user.NewController(*userService)
	return &Server{
		userController: *userController,
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, context echo.Context) {
		if errors.Is(err, user.ErrInvalidPassword) {
			context.JSON(http.StatusBadRequest, map[string]string{"message": "invalid password"})
			return
		}
		e.DefaultHTTPErrorHandler(err, context)
	}
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteUser(g, s.userController)
}

func RouteUser(e *echo.Group, c user.Controller) {
	e.POST("/login", c.Login)
}
