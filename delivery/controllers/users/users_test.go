package users

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"project-airbnb/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

func TestUser(t *testing.T) {
	t.Run("Test Error Request Register", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"name":     1,
			"email":    1,
			"password": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/userss")

		userController := NewUsersControllers(mockFalseUserRepository{})
		userController.RegisterUserCtrl()(context)

		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Register", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/userss")

		userController := NewUsersControllers(mockFalseUserRepository{})
		userController.RegisterUserCtrl()(context)

		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Test Register", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})
		userController.RegisterUserCtrl()(context)

		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
		assert.Equal(t, "ilham", response.Data.Name)
		assert.Equal(t, "ilham@gmail.com", response.Data.Email)
	})

	// LOGIN
	t.Run("Test Error Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@gmail.com",
			"password": "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := NewUsersControllers(mockFalseUserRepository{})
		authControl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Not Found", responses.Message)
	})
	t.Run("Test Error Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]int{
			"email":    1,
			"password": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := NewUsersControllers(mockFalseUserRepository{})
		authControl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Bad Request", responses.Message)
	})
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@gmail.com",
			"password": "ilham123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := NewUsersControllers(mockUserRepository{})
		authControl.LoginAuthCtrl()(context)

		responses := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Error Get All User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})

		userController.GetUsersCtrl()(context)
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Test Get All User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})

		userController.GetUsersCtrl()(context)
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test GetByID User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(userController.GetUserByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error GetByID User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(userController.GetUserByIdCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Error Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Error Request Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"name":     1,
			"email":    1,
			"password": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Error Delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TestingUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
}

type mockUserRepository struct{}

func (ma mockUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, nil
}
func (m mockUserRepository) Gets() ([]entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return []entities.User{
		{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString},
	}, nil
}
func (m mockUserRepository) Get(userid int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, nil
}

func (m mockUserRepository) Register(user entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, nil
}

func (m mockUserRepository) Update(user entities.User, id int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, nil
}

func (m mockUserRepository) Delete(id int) (entities.User, error) {
	return entities.User{ID: 1, Name: "Ilham", Email: "ilham@gmail.com"}, nil
}

type mockFalseUserRepository struct{}

func (ma mockFalseUserRepository) LoginUser(email, password string) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, errors.New("False Login Object")
}
func (m mockFalseUserRepository) Gets() ([]entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return []entities.User{
		{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString},
	}, errors.New("False Login Object")
}
func (m mockFalseUserRepository) Get(userid int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, errors.New("False Login Object")
}

func (m mockFalseUserRepository) Register(user entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, errors.New("False Login Object")
}

func (m mockFalseUserRepository) Update(user entities.User, id int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	passwordString := fmt.Sprintf("%x", hash[:])
	return entities.User{ID: 1, Name: "ilham", Email: "ilham@gmail.com", Password: passwordString}, errors.New("False Login Object")
}

func (m mockFalseUserRepository) Delete(id int) (entities.User, error) {
	return entities.User{ID: 1, Name: "Ilham", Email: "ilham@gmail.com"}, errors.New("False Login Object")
}
