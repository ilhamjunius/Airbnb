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
	e.DELETE("/users", uctrl.DeleteUserCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.PUT("/users", uctrl.UpdateUserCtrl(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/user", uctrl.GetUserByIdCtrl(), middleware.JWT([]byte("RAHASIA")))
	// ---------------------------------------------------------------------
	// CRUD Rooms
	// ---------------------------------------------------------------------
	e.POST("/rooms", rmCtrl.Create(), middleware.JWT([]byte("RAHASIA")))

	e.GET("/myrooms", rmCtrl.Gets(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/rooms", rmCtrl.Get(), middleware.JWT([]byte("RAHASIA")))
	e.PUT("/rooms/:id", rmCtrl.Update(), middleware.JWT([]byte("RAHASIA")))
	e.DELETE("/rooms/:id", rmCtrl.Delete(), middleware.JWT([]byte("RAHASIA")))
	// ---------------------------------------------------------------------
	// CRUD Transactions
	// ---------------------------------------------------------------------
	e.GET("/transactions", trctrl.Gets(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/transactions/order", trctrl.Get(), middleware.JWT([]byte("RAHASIA")))
	e.POST("/transactions", trctrl.Update())

	// ---------------------------------------------------------------------
	// CRUD Bookings
	// ---------------------------------------------------------------------
	e.POST("/booking", bkCtrl.Create(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/booking", bkCtrl.Get(), middleware.JWT([]byte("RAHASIA")))
	e.GET("/bookings", bkCtrl.Gets(), middleware.JWT([]byte("RAHASIA")))

}
