package transactions

import (
	"fmt"
	"net/http"
	"project-airbnb/delivery/common"
	"project-airbnb/repository/transactions"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	// "github.com/midtrans/midtrans-go/coreapi"
)

type TransactionsController struct {
	Repo transactions.TransactionsInterface
}

func NewTransactionsControllers(tsrep transactions.TransactionsInterface) *TransactionsController {
	return &TransactionsController{Repo: tsrep}
}

func (trrep TransactionsController) Gets() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if res, err := trrep.Repo.Gets(uint(userID)); err != nil || len(res) == 0 {
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
func (trrep TransactionsController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		if res, err := trrep.Repo.Get(uint(userID)); err != nil || len(res) == 0 {
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

//manual parse
func (trrep TransactionsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		updateRoom := UpdateTransactionsRequestFormat{}
		if err := c.Bind(&updateRoom); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if res, err := trrep.Repo.Update(updateRoom.InvoiceID, "settlement"); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			responses := UpdateTransactionsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    res,
			}
			return c.JSON(http.StatusOK, responses)
		}
	}
}

// var crc coreapi.Client

func (trrep TransactionsController) UpdateCallBack() echo.HandlerFunc {
	return func(c echo.Context) error {

		midtrans.ServerKey = "SB-Mid-server-W-ANVsQXp9S7q65qndszXrcD"
		midtrans.ClientKey = "SB-Mid-client-QVIZg4p30WL2WLy8"
		midtrans.Environment = midtrans.Sandbox

		var notificationPayload map[string]interface{}
		if err := c.Bind(&notificationPayload); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		// orderID, exists := notificationPayload["order-id"].(string)
		// if !exists {
		// 	fmt.Println("not found")
		// }
		fmt.Println("notification", notificationPayload)
		fmt.Println("len", len(notificationPayload))
		fmt.Println("invoice", notificationPayload["order_id"].(string))
		fmt.Println("status", notificationPayload["transaction_status"].(string))
		if res, err := trrep.Repo.Update(notificationPayload["order_id"].(string), notificationPayload["transaction_status"].(string)); err != nil || res.ID == 0 {
			fmt.Println("intip res", res)
			fmt.Println("intip err", err)
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}
		// i:=0
		// if i!=len(notificationPayload){

		// }

		// transactionStatusResp, err := crc.CheckTransaction(orderID)
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		// }
		// if transactionStatusResp != nil {
		//
		// }

	}
}
