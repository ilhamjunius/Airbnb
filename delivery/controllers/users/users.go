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
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		token, err := CreateTokenAuth(checkedUser.ID)
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successful Operation",
			"token":   token,
		},
		)
	}
}
func (uscon UsersController) RegisterUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := RegisterUserRequestFormat{}
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
		fmt.Println(newUser.Email)
		res, err := uscon.Repo.Register(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		data := UserResponse{
			ID:    res.ID,
			Name:  res.Name,
			Email: res.Email,
		}
		response := RegisterUserResponseFormat{
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
			response := GetUsersResponseFormat{
				Code:    http.StatusOK,
				Message: "Successful Opration",
				Data:    users,
			}
			return c.JSON(http.StatusOK, response)
		}

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
