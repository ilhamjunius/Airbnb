package users

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"project-airbnb/delivery/common"
	"project-airbnb/entities"
	"project-airbnb/repository/users"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo users.UsersInterface
}

func NewUsersControllers(usrep users.UsersInterface) *UsersController {
	return &UsersController{Repo: usrep}
}
func (uscon UsersController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(loginFormat.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		checkedUser, err := uscon.Repo.LoginUser(loginFormat.Email, stringPassword)
		if err != nil || checkedUser.Email == "" {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		token, _ := CreateTokenAuth(checkedUser.ID)

		return c.JSON(http.StatusOK, LoginResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Token:   token,
		})

	}
}
func (uscon UsersController) RegisterUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := UserRequestFormat{}
		if err := c.Bind(&newUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(newUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		newUser := entities.User{
			Email:    newUserReq.Email,
			Password: stringPassword,
			Name:     newUserReq.Name,
		}
		res, err := uscon.Repo.Register(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := UserResponse{
			ID:    res.ID,
			Name:  res.Name,
			Email: res.Email,
		}
		response := UserResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}

		return c.JSON(http.StatusOK, response)
	}
}
func (uscon UsersController) GetUserByIdCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		user, err := uscon.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		data := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		response := UserResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}

		return c.JSON(http.StatusOK, response)
	}
}
func (uscon UsersController) GetUsersCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		if users, err := uscon.Repo.Gets(); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {
			data := []UserResponse{}
			for _, user := range users {
				data = append(
					data, UserResponse{
						ID:    user.ID,
						Name:  user.Name,
						Email: user.Email,
					},
				)
			}
			response := GetUsersResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Operation",
				Data:    data,
			}
			return c.JSON(http.StatusOK, response)
		}

	}
}
func (uscon UsersController) DeleteUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		deletedUser, err := uscon.Repo.Delete(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		data := UserResponse{
			ID:    deletedUser.ID,
			Name:  deletedUser.Name,
			Email: deletedUser.Email,
		}
		response := UserResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}

		return c.JSON(http.StatusOK, response)
	}
}
func (uscon UsersController) UpdateUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		updateUserReq := UserRequestFormat{}
		if err := c.Bind(&updateUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(updateUserReq.Password))
		stringPassword := fmt.Sprintf("%x", hash[:])
		updateUser := entities.User{
			Email: updateUserReq.Email,
			Name:  updateUserReq.Name,
		}
		if updateUserReq.Password != "" {
			updateUser.Password = stringPassword
		}
		res, err := uscon.Repo.Update(updateUser, id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		data := UserResponse{
			ID:    res.ID,
			Name:  res.Name,
			Email: res.Email,
		}
		response := UserResponseFormat{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    data,
		}

		return c.JSON(http.StatusOK, response)
	}
}
func CreateTokenAuth(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RAHASIA"))
}
