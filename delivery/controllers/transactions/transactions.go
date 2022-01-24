package transactions

import (
	"net/http"
	"project-airbnb/delivery/common"
	"project-airbnb/repository/transactions"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionsController struct {
	Repo transactions.TransactionsInterface
}

func NewTransactionsControllers(tsrep transactions.TransactionsInterface) *TransactionsController {
	return &TransactionsController{Repo: tsrep}
}

func (trrep TransactionsController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if res, err := trrep.Repo.Get(uint(userID)); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"code":    200,
				"message": "Successful Operation",
				"data":    res,
			})
		}

	}
}
