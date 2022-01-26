package transactions

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"project-airbnb/delivery/controllers/users"
	"project-airbnb/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

// var crc coreapi.Client

func TestTransactionTrue(t *testing.T) {
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
	t.Run("Test Get transactions", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		transactionController := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response["message"])
	})
	t.Run("Test Get transactions", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		transactionController := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response["message"])
	})
	t.Run("Test Update transactions", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"invoice_id": "invoice",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tramsactions")

		transactionController := NewTransactionsControllers(mockTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response["message"])
	})
	t.Run("Test Update Callback transactions", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"order_id":           "INV-2/book/14d87b52-75de-47e0-8e19-0322dc4149b3",
			"transaction_status": "settlement",
		})
		req := httptest.NewRequest(http.MethodPost, "/transactions/callback", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/transactions/callback")

		transactionController := NewTransactionsControllers(mockTransactionRepository{})
		transactionController.UpdateCallBack()(context)

		response := users.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		fmt.Println("Response:", response)
		assert.Equal(t, "Successful Operation", "")
	})

}

type mockTransactionRepository struct{}

func (m mockTransactionRepository) Get(userID uint) ([]entities.Transaction, error) {
	return []entities.Transaction{
		{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "Paid"},
	}, nil
}
func (m mockTransactionRepository) Gets(userID uint) ([]entities.Transaction, error) {
	return []entities.Transaction{
		{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "Paid"},
	}, nil
}
func (m mockTransactionRepository) Update(invoiceID, status string) (entities.Transaction, error) {
	return entities.Transaction{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "Paid"}, nil
}

// func (m mockTransactionRepository) UpdateCallBack(invoiceID, status string) (entities.Transaction, error) {
// 	return entities.Transaction{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "settlement"}, nil
// }

func TestTransactionFalse(t *testing.T) {
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
	t.Run("Test Error Get transactions", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		transactionController := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response["message"])
	})
	t.Run("Test Error Get transactions", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		transactionController := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response["message"])
	})
	t.Run("Test Error Update transactions", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"invoice_id": "invoice",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		transactionController := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response["message"])
	})
	t.Run("Test Error Update Request transactions", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"invoice_id": 1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/transaction")

		transactionController := NewTransactionsControllers(mockFalseTransactionRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(transactionController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response map[string]interface{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response["message"])
	})

}

type mockFalseTransactionRepository struct{}

func (m mockFalseTransactionRepository) Get(userID uint) ([]entities.Transaction, error) {
	return []entities.Transaction{
		{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "Paid"},
	}, errors.New("False Login Object")
}
func (m mockFalseTransactionRepository) Gets(userID uint) ([]entities.Transaction, error) {
	return []entities.Transaction{
		{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "Paid"},
	}, errors.New("False Login Object")
}
func (m mockFalseTransactionRepository) Update(invoiceID, status string) (entities.Transaction, error) {
	return entities.Transaction{ID: 1, Invoice: "invoice", Url: "https://app.sandbox.midtrans.com/snap/v2/vtweb/56cfa145-ccb0-4e7a-a79c-450014f4edb3", Status: "Paid"}, errors.New("False Login Object")
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
