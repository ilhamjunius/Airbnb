package rooms

import (
	"net/http"
	"project-airbnb/delivery/common"
	"project-airbnb/entities"
	"project-airbnb/repository/rooms"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RoomsController struct {
	Repo rooms.RoomsInterface
}

func NewRoomsControllers(ri rooms.RoomsInterface) *RoomsController {
	return &RoomsController{Repo: ri}
}

func (rrcon RoomsController) Gets() echo.HandlerFunc {
	return func(c echo.Context) error {
		if res, err := rrcon.Repo.Gets(); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			response := GetRoomsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    res,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
}

func (rrcon RoomsController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		newRoomReq := AddNewRoomRequestFormat{}

		if err := c.Bind(&newRoomReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newRoom := entities.Room{
			User_id:  uint(userID),
			Name:     newRoomReq.Name,
			Location: newRoomReq.Location,
			Price:    newRoomReq.Price,
		}

		if res, err := rrcon.Repo.Create(newRoom); err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}

	}
}
func (rrcon RoomsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		newRoomReq := AddNewRoomRequestFormat{}

		if err := c.Bind(&newRoomReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newRoom := entities.Room{
			User_id:  uint(userID),
			Name:     newRoomReq.Name,
			Location: newRoomReq.Location,
			Price:    newRoomReq.Price,
		}

		if res, err := rrcon.Repo.Update(newRoom); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			response := GetRoomsResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    res,
			}
			return c.JSON(http.StatusOK, response)
		}

	}
}
func (rrcon RoomsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))

		delRoomReq := DeleteRoomRequestFormat{}
		if err := c.Bind(&delRoomReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		if res, err := rrcon.Repo.Delete(delRoomReq.RoomID, uint(userID)); err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		} else {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}
	}
}
