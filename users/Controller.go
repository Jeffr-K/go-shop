package user

import (
	"github.com/labstack/echo/v4"
	"github.com/lambsroad/go-shop/users/dto"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller Controller) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	res, err := controller.service.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller Controller) Login(c echo.Context) error {
	userName := c.FormValue("user_name")
	password := c.FormValue("password")

	res, err := controller.service.Login(userName, password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
