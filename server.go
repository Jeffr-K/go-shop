package go_shop

import (
	"fmt"
	"github.com/labstack/echo/v4"
	user "github.com/lambsroad/go-shop/users"
	"log"
)

const _defaultPort = 8080

type Server struct {
	userController user.Controller
}

func NewDefaultServer() *Server {
	return &Server{
		userController: *user.NewController(*user.NewService(user.DefaultSecret)),
	}
}

func (s *Server) Run() {
	e := echo.New()
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteUser(g, s.userController)
}

func RouteUser(e *echo.Group, c user.Controller) {
	e.POST("/login", c.Login)
}
