package handler

import (
	"net/http"
	"strings"
	"wanderer/features/users"

	echo "github.com/labstack/echo/v4"
)

func NewUserHandler(userService users.Service) users.Handler {
	return &userHandler{
		userService: userService,
	}
}

type userHandler struct {
	userService users.Service
}

func (hdl *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(RegisterRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "incorrect input data"
			return c.JSON(http.StatusBadRequest, response)
		}

		var data = request.ToEntity()

		if err := hdl.userService.Register(*data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "Duplicate") {
				response["message"] = "email is already in use"
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "register success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *userHandler) Login() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}
