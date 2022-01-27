package books

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"project-airbnb/configs"
	"project-airbnb/delivery/controllers/users"
	"project-airbnb/entities"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config := configs.GetConfig()
	fmt.Print(config)
	m.Run()

}

var jwtToken string

func TestBookingTrue(t *testing.T) {
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

		userControl := users.NewUsersControllers(mockUserRepository{})
		userControl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Get All Booking", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/bookings")

		bookController := NewBooksControllers(mockBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Get All Booking", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/bookings")

		bookController := NewBooksControllers(mockBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Insert Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"room_id": 1,
			"price":   500000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("Test Update Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"room_id":  1,
			"duration": 5,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
}

type mockBookRepository struct{}

func (m mockBookRepository) Gets(userID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, nil
}
func (m mockBookRepository) Get(userID, roomID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, nil
}
func (m mockBookRepository) Create(newBooking entities.Book) (entities.Book, error) {
	return entities.Book{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1}, nil
}
func (m mockBookRepository) Update(bookID uint) (entities.Book, error) {
	return entities.Book{ID: 1, User_id: 2, Room_id: 1, Checkin: time.Time{}, Checkout: time.Time{}, Transaction_id: 1}, nil
}
func (m mockBookRepository) CreateTransactions(userID, roomID uint, invoiceID string, duration int) (entities.Transaction, error) {
	return entities.Transaction{ID: 1, Invoice: "INV-3/book/641a10e2-344b-4021-b23c-7035821853ec", Status: "PENDING", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/ecf407a9-814a-46b8-afc6-aa82b67c3496"}, nil
}

func TestBookingFalse(t *testing.T) {
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

		userControl := users.NewUsersControllers(mockUserRepository{})
		userControl.LoginAuthCtrl()(context)

		responses := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Error Get All Booking", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/bookings")

		bookController := NewBooksControllers(mockFalseBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Accepted", response.Message)
	})
	t.Run("Test Error Request Get All Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"room_id": "abc",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/bookings")

		bookController := NewBooksControllers(mockFalseBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Gets All Booking", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/bookings")

		bookController := NewBooksControllers(mockFalseBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Error Insert Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"room_id": 1,
			"price":   500000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockFalseBookRepository1{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Test Error Insert Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"room_id": 1,
			"price":   500000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockFalseCreateBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Error Request Insert Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"room_id": "abc",
			"price":   "500000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockFalseBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Request Update Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"room_id":  "abc",
			"duration": 5,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockFalseBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Request Update Booking", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"room_id": 0,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/booking")

		bookController := NewBooksControllers(mockFalseBookRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(bookController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := TransactionsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
}

type mockFalseBookRepository struct{}

func (m mockFalseBookRepository) Gets(userID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, errors.New("False Login Object")
}
func (m mockFalseBookRepository) Get(userID, roomID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, errors.New("False Login Object")
}
func (m mockFalseBookRepository) Create(newBooking entities.Book) (entities.Book, error) {
	return entities.Book{ID: 0, User_id: 2, Room_id: 1, Transaction_id: 1}, errors.New("False Login Object")
}
func (m mockFalseBookRepository) Update(BookID uint) (entities.Book, error) {
	return entities.Book{ID: 0, Room_id: 0}, errors.New("Not Found")
}
func (m mockFalseBookRepository) CreateTransactions(userID, roomID uint, invoiceID string, duration int) (entities.Transaction, error) {
	return entities.Transaction{ID: 0, Invoice: "", Status: "PENDING", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/ecf407a9-814a-46b8-afc6-aa82b67c3496"}, nil
}

type mockFalseBookRepository1 struct{}

func (m mockFalseBookRepository1) Gets(userID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, errors.New("False Login Object")
}
func (m mockFalseBookRepository1) Get(userID, roomID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, errors.New("False Login Object")
}
func (m mockFalseBookRepository1) Create(newBooking entities.Book) (entities.Book, error) {
	return entities.Book{ID: 0, User_id: 2, Room_id: 1, Transaction_id: 1}, errors.New("False Login Object")
}
func (m mockFalseBookRepository1) Update(bookID uint) (entities.Book, error) {
	return entities.Book{ID: 0, Room_id: 0}, errors.New("Not Found")
}
func (m mockFalseBookRepository1) CreateTransactions(userID, roomID uint, invoiceID string, duration int) (entities.Transaction, error) {
	return entities.Transaction{ID: 1, Invoice: "", Status: "PENDING", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/ecf407a9-814a-46b8-afc6-aa82b67c3496"}, nil
}

type mockFalseCreateBookRepository struct{}

func (m mockFalseCreateBookRepository) Gets(userID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, errors.New("False Login Object")
}
func (m mockFalseCreateBookRepository) Get(userID, roomID uint) ([]entities.Book, error) {
	return []entities.Book{
		{ID: 1, User_id: 2, Room_id: 1, Transaction_id: 1},
	}, errors.New("False Login Object")
}
func (m mockFalseCreateBookRepository) Create(newBooking entities.Book) (entities.Book, error) {
	return entities.Book{ID: 0, User_id: 2, Room_id: 1, Transaction_id: 1}, errors.New("False Login Object")
}
func (m mockFalseCreateBookRepository) Update(bookID uint) (entities.Book, error) {
	return entities.Book{ID: 1, User_id: 2, Room_id: 1, Checkin: time.Time{}, Checkout: time.Time{}, Transaction_id: 1}, errors.New("False Login Object")
}
func (m mockFalseCreateBookRepository) CreateTransactions(userID, roomID uint, invoiceID string, duration int) (entities.Transaction, error) {
	return entities.Transaction{ID: 0, Invoice: "", Status: "PENDING", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/ecf407a9-814a-46b8-afc6-aa82b67c3496"}, errors.New("False Login Object")
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
