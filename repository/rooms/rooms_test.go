package rooms

import (
	"crypto/sha256"
	"fmt"
	"project-airbnb/configs"
	"project-airbnb/entities"
	"project-airbnb/repository/books"
	"project-airbnb/repository/users"
	"project-airbnb/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	m.Run()
	config := configs.GetConfig()
	db := utils.InitDB(config)
	// db.Migrator().DropTable(&entities.User{})
	// db.Migrator().DropTable(&entities.Room{})
	// db.Migrator().DropTable(&entities.Book{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Book{})

}
func TestRoomsRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.User{})
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.Book{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Room{})
	db.AutoMigrate(&entities.Book{})
	roomRepo := NewRoomsRepo(db)
	userRepo := users.NewUsersRepo(db)
	bookRepo := books.NewBooksRepo(db)

	db.Migrator().DropTable(&entities.Room{})
	t.Run("Get All My Room", func(t *testing.T) {

		_, err := roomRepo.Gets(100)
		assert.Error(t, err)

	})
	t.Run("Get All My Room", func(t *testing.T) {

		_, err := roomRepo.GetsById(100, 100)
		assert.Error(t, err)

	})
	t.Run("Get All Room", func(t *testing.T) {

		_, err := roomRepo.Get(100)
		assert.Error(t, err)

	})
	t.Run("Get All Room", func(t *testing.T) {

		_, err := roomRepo.GetById(100, 100)
		assert.Error(t, err)

	})
	t.Run("Get All Room Sold", func(t *testing.T) {

		_, err := roomRepo.GetMyRoomIncome(100)
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.Room{})
	t.Run("Register", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilham123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "ilham"

		_, err := userRepo.Register(mockUser)
		assert.Nil(t, err)

	})
	t.Run("Register", func(t *testing.T) {
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
		mockRoom.Name = "Room2"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 300000
		mockRoom.Duration = 7
		mockRoom.User_id = 2
		mockRoom.Status = "OPEN"

		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Update Room into Database", func(t *testing.T) {

		var mockRoom entities.Room
		mockRoom.Name = "Room2"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 300000
		mockRoom.Duration = 6
		mockRoom.User_id = 1
		mockRoom.Status = "OPEN"

		_, err := roomRepo.Update(mockRoom, 1)
		assert.Nil(t, err)

	})
	t.Run("Insert Transaction into Database", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 1
		invoice := "INV-2/book/41a74c38-2880-4d91-8875-f8f0f06a641c"
		_, err := bookRepo.CreateTransactions(mockBook.User_id, mockBook.Room_id, invoice, 2)
		assert.Nil(t, err)

	})
	t.Run("Insert Booking into Database", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 1
		mockBook.Transaction_id = 1

		_, err := bookRepo.Create(mockBook)
		assert.Nil(t, err)

	})
	t.Run("Get All Room Sold", func(t *testing.T) {

		_, err := roomRepo.GetMyRoomIncome(1)
		assert.Nil(t, err)

	})
	t.Run("Get All Room", func(t *testing.T) {

		_, err := roomRepo.Gets(1)
		assert.Nil(t, err)

	})
	t.Run("Get All Room By Room Id", func(t *testing.T) {

		_, err := roomRepo.GetsById(2, 2)
		assert.Nil(t, err)

	})
	t.Run("Get All Room into Database", func(t *testing.T) {

		_, err := roomRepo.Get(100)
		assert.Nil(t, err)

	})
	t.Run("Get All Room By Id", func(t *testing.T) {

		_, err := roomRepo.GetById(1, 2)
		assert.Nil(t, err)

	})
	t.Run("Delete Room into Database", func(t *testing.T) {

		_, err := roomRepo.Delete(1, 1)
		assert.Nil(t, err)

	})
}
