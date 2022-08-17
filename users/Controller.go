package user

import (
	"github.com/labstack/echo/v4"
	"github.com/lambsroad/go-shop/users/dto"
	"net/http"
)

type Controller struct {
	Service Service
}

func NewController(service Service) *Controller {
	return &Controller{Service: service}
}

func (controller Controller) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	res, err := controller.Service.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller Controller) Login(c echo.Context) error {
	userName := c.FormValue("user_name")
	password := c.FormValue("password")

	res, err := controller.Service.Login(userName, password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (controller Controller) FindUserById(c echo.Context) error {
	userId := c.Param("user_id")
	res, err := controller.Service.FindUserOneByID(userId)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
