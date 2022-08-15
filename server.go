package go_shop

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	user "github.com/lambsroad/go-shop/users"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"net/http"
)

const _defaultPort = 4000

type Server struct {
	userController user.Controller
}

func NewDefaultServer() *Server {
	return &Server{
		userController: *user.NewController(*user.NewService(user.UserRepository{}))}
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
	g.GET("/swagger/*", echoSwagger.WrapHandler)
	RouteUser(g, s.userController)
}

func RouteUser(e *echo.Group, c user.Controller) {
	e.POST("/create", c.CreateUser)
	e.POST("/login", c.Login)
	e.GET("/user_id", c.FindUserById)
}
