package users

import (
	"net/http"
	"project-airbnb/delivery/common"
	"project-airbnb/repository/users"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo users.UsersInterface
}

func NewUsersControllers(usrep users.UsersInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

func (uscon UsersController) GetUsersCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		if users, err := uscon.Repo.Gets(); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			response := GetUsersResponseFormat{
				Message: "Successful Opration",
				Data:    users,
			}
			return c.JSON(http.StatusOK, response)
		}

	}
}
