package routes

import (
	"project-airbnb/delivery/controllers/users"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController) {
	e.GET("/users", uctrl.GetUsersCtrl())
	e.POST("/login", uctrl.LoginAuthCtrl())
	e.POST("/register", uctrl.RegisterUserCtrl())
}
