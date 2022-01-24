package routes

import (
	"project-airbnb/delivery/controllers/books"
	"project-airbnb/delivery/controllers/rooms"
	"project-airbnb/delivery/controllers/transactions"
	"project-airbnb/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uctrl *users.UsersController, rmCtrl *rooms.RoomsController, bkCtrl *books.BooksController, trctrl *transactions.TransactionsController) {
	// ---------------------------------------------------------------------
	// CRUD Users
	// ---------------------------------------------------------------------
	e.GET("/users", uctrl.GetUsersCtrl())
	e.POST("/login", uctrl.LoginAuthCtrl())
	e.POST("/register", uctrl.RegisterUserCtrl())

	// ---------------------------------------------------------------------
	// CRUD Rooms
	// ---------------------------------------------------------------------
	e.POST("/rooms", rmCtrl.Create(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/rooms", rmCtrl.Gets())
	e.PUT("/rooms", rmCtrl.Update(), middleware.JWT([]byte("RAHASIA")))
	e.DELETE("/rooms", rmCtrl.Delete(), middleware.JWT([]byte("RAHASIA")))
	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.GET("/transactions", trctrl.Get(), middleware.JWT([]byte("RAHASIA")))
	e.POST("/booking", bkCtrl.Create(), middleware.JWT([]byte("RAHASIA")))

}
