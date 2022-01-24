package routes

import (
	"project-airbnb/delivery/controllers/transactions"
	"project-airbnb/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController, trctrl *transactions.TransactionsController) {
	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.GET("/users", uctrl.GetUsersCtrl())

	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.GET("/transactions", trctrl.Get(), middleware.JWT([]byte("RAHASIA")))
}
