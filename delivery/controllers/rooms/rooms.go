package rooms

import (
	"net/http"
	"project-airbnb/delivery/common"
	"project-airbnb/entities"
	"project-airbnb/repository/rooms"
	"strconv"

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
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		rooms, err := rrcon.Repo.Gets(userID)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		data := []RoomResponse{}

		for _, room := range rooms {
			data = append(
				data, RoomResponse{
					ID:       room.ID,
					Name:     room.Name,
					Location: room.Location,
					Duration: room.Duration,
					User_id:  room.User_id,
					Price:    room.Price,
					Status:   room.Status,
				},
			)
		}
		response := GetRoomsResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

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
			Duration: newRoomReq.Duration,
			Status:   newRoomReq.Status,
		}
		res, err := rrcon.Repo.Create(newRoom)
		if err != nil || res.ID == 0 {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := RoomResponse{
			ID:       res.ID,
			Name:     res.Name,
			Location: res.Location,
			Duration: res.Duration,
			User_id:  res.User_id,
			Price:    res.Price,
			Status:   res.Status,
		}
		response := GetRoomsResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (rrcon RoomsController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
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
			Duration: newRoomReq.Duration,
			Status:   newRoomReq.Status,
		}
		res, err := rrcon.Repo.Update(newRoom, roomId)
		if err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		data := RoomResponse{
			ID:       res.ID,
			Name:     res.Name,
			Location: res.Location,
			Duration: res.Duration,
			User_id:  res.User_id,
			Price:    res.Price,
			Status:   res.Status,
		}
		response := GetRoomsResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
func (rrcon RoomsController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		roomId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := int(claims["userid"].(float64))
		res, err := rrcon.Repo.Delete(roomId, uint(userID))
		if err != nil || res.ID == 0 {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		data := RoomResponse{
			ID:       res.ID,
			Name:     res.Name,
			Location: res.Location,
			Duration: res.Duration,
			User_id:  res.User_id,
			Price:    res.Price,
			Status:   res.Status,
		}
		response := GetRoomsResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}
		return c.JSON(http.StatusOK, response)

	}
}
