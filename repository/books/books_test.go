package books

import (
	"crypto/sha256"
	"fmt"
	"project-airbnb/configs"
	"project-airbnb/entities"
	"project-airbnb/repository/rooms"
	"project-airbnb/repository/users"
	"project-airbnb/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	config := configs.GetConfig()
	fmt.Println(config)
	m.Run()
}
func TestBookingRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	bookingRepo := NewBooksRepo(db)
	userRepo := users.NewUsersRepo(db)
	roomRepo := rooms.NewRoomsRepo(db)
	t.Run("Insert User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilham123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "ilham"

		_, err := userRepo.Register(mockUser)
		assert.Nil(t, err)

	})
	t.Run("Insert User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("junius123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "junius@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "junius"

		_, err := userRepo.Register(mockUser)
		assert.Nil(t, err)

	})
	t.Run("Insert Room into Database", func(t *testing.T) {
		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 300000
		mockRoom.Duration = 7
		mockRoom.User_id = 1
		mockRoom.Status = "OPEN"

		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Insert Room into Database", func(t *testing.T) {
		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 300000
		mockRoom.Duration = 7
		mockRoom.User_id = 1
		mockRoom.Status = "CLOSED"

		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Insert Transaction into Database", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 1
		invoice := "INV-2/book/41a74c38-2880-4d91-8875-f8f0f06a641c"
		_, err := bookingRepo.CreateTransactions(mockBook.User_id, mockBook.Room_id, invoice, 0)
		assert.Nil(t, err)

	})
	t.Run("Insert Transaction into Database", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 2
		invoice := "INV-2/book/41a74c38-2880-4d91-8875-f8f0f06a641c"
		_, err := bookingRepo.CreateTransactions(mockBook.User_id, mockBook.Room_id, invoice, 0)
		assert.Nil(t, err)

	})
	t.Run("Insert Booking into Database", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 1
		mockBook.Transaction_id = 1

		_, err := bookingRepo.Create(mockBook)
		assert.Nil(t, err)

	})
	t.Run("Insert Booking into Database", func(t *testing.T) {
		_, err := bookingRepo.Get(2, 1)
		assert.Nil(t, err)

	})
	t.Run("Show Booking into Database", func(t *testing.T) {
		_, err := bookingRepo.Gets(2)
		assert.Nil(t, err)

	})
	t.Run("Update booking into Database", func(t *testing.T) {
		res, err := bookingRepo.Update(1, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
