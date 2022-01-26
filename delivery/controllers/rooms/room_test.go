package rooms

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
	"project-airbnb/delivery/common"
	"project-airbnb/delivery/controllers/users"
	"project-airbnb/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

func TestMain(m *testing.M) {
	config := configs.GetConfig()
	fmt.Print(config)
	m.Run()

}
func TestUserTrue(t *testing.T) {

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
	t.Run("Test Insert Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"id":       1,
			"name":     "Room1",
			"location": "Bandung",
			"user_id":  1,
			"price":    300000,
			"duration": 7,
			"status":   "Already Booked",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms")

		roomController := NewRoomsControllers(mockRoomRepository{})
		if err := middleware.JWT([]byte(common.JWT_SECRET_KEY))(roomController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Update Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"id":       1,
			"name":     "Room1",
			"location": "Bandung",
			"user_id":  1,
			"price":    350000,
			"duration": 7,
			"status":   "Already Booked",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		roomController := NewRoomsControllers(mockRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Get My Room", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/room")

		roomController := NewRoomsControllers(mockRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Get All Room", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/rooms")

		roomController := NewRoomsControllers(mockRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Delete Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"room_id": 1,
		})
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms")
		context.SetParamNames("id")
		context.SetParamValues("1")

		roomController := NewRoomsControllers(mockRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})

}

type mockRoomRepository struct{}

func (m mockRoomRepository) Gets(userId int) ([]entities.Room, error) {
	return []entities.Room{
		{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"},
	}, nil
}
func (m mockRoomRepository) GetsById(userId, roomId int) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked", Desciption: "Deskripsi Room 1"}, nil
}
func (m mockRoomRepository) Get(userId int) ([]entities.Room, error) {
	return []entities.Room{
		{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"},
	}, nil
}
func (m mockRoomRepository) GetById(userId, roomId int) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked", Desciption: "Deskripsi Room 1"}, nil
}
func (m mockRoomRepository) Create(newRoom entities.Room) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"}, nil
}
func (m mockRoomRepository) Update(editRoom entities.Room, roomId int) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"}, nil
}
func (m mockRoomRepository) Delete(roomID int, userID uint) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"}, nil
}

func TestUserFalse(t *testing.T) {
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
	t.Run("Test Error Insert Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"id":       1,
			"name":     "Room1",
			"location": "Bandung",
			"user_id":  1,
			"price":    300000,
			"duration": 7,
			"status":   "Already Booked",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms")

		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Test Error Request Insert Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"id":       "1",
			"name":     "Room1",
			"location": "Bandung",
			"user_id":  "1",
			"price":    "300000",
			"duration": "7",
			"status":   "Already Booked",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms")

		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Update Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"id":       1,
			"name":     "Room1",
			"location": "Bandung",
			"user_id":  1,
			"price":    350000,
			"duration": 7,
			"status":   "Already Booked",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Error Request Update Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"id":       "1",
			"name":     "Room1",
			"location": "Bandung",
			"user_id":  "1",
			"price":    "350000",
			"duration": "7",
			"status":   "Already Booked",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms")

		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Request Update Room", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"id":       1,
			"name":     1,
			"location": 1,
			"user_id":  1,
			"price":    1,
			"duration": 1,
			"status":   1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms")
		context.SetParamNames("id")
		context.SetParamValues("1")
		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Get My Room", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/room")

		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Error Get All Room", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/rooms")

		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Error  Request Delete Room", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")
		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Test Error Delete Room", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/rooms/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		roomController := NewRoomsControllers(mockFalseRoomRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(roomController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRoomsResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})

}

type mockFalseRoomRepository struct{}

func (m mockFalseRoomRepository) Gets(userId int) ([]entities.Room, error) {
	return []entities.Room{
		{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"},
	}, errors.New("False Login Object")
}
func (m mockFalseRoomRepository) GetsById(userId, roomId int) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked", Desciption: "Deskripsi Room 1"}, errors.New("False Login Object")
}
func (m mockFalseRoomRepository) Get(userId int) ([]entities.Room, error) {
	return []entities.Room{
		{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"},
	}, errors.New("False Login Object")
}
func (m mockFalseRoomRepository) GetById(userId, roomId int) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked", Desciption: "Deskripsi Room 1"}, errors.New("False Login Object")
}
func (m mockFalseRoomRepository) Create(newRoom entities.Room) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"}, errors.New("False Login Object")
}
func (m mockFalseRoomRepository) Update(editRoom entities.Room, roomId int) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"}, errors.New("False Login Object")
}
func (m mockFalseRoomRepository) Delete(roomID int, userID uint) (entities.Room, error) {
	return entities.Room{ID: 1, Name: "Room1", User_id: 1, Location: "Bandung", Price: 500000, Duration: 7, Status: "Already Booked"}, errors.New("False Login Object")
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
