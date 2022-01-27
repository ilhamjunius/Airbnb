package transactions

import (
	"crypto/sha256"
	"fmt"
	"project-airbnb/configs"
	"project-airbnb/entities"
	"project-airbnb/repository/books"
	"project-airbnb/repository/rooms"
	"project-airbnb/repository/users"
	"project-airbnb/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionsRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	transactionRepo := NewTransactionsRepo(db)
	bookingRepo := books.NewBooksRepo(db)
	userRepo := users.NewUsersRepo(db)
	roomRepo := rooms.NewRoomsRepo(db)
	t.Run("Insert User_id 1", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilham123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "ilham"

		_, err := userRepo.Register(mockUser)
		assert.Nil(t, err)

	})
	t.Run("Insert User_id 2", func(t *testing.T) {
		hash := sha256.Sum256([]byte("junius123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "junius@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "junius"

		_, err := userRepo.Register(mockUser)
		assert.Nil(t, err)

	})
	t.Run("Insert Room_id 1", func(t *testing.T) {
		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 10000
		mockRoom.Duration = 10
		mockRoom.User_id = 1
		mockRoom.Status = "OPEN"
		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Insert Room_id 2", func(t *testing.T) {
		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 20000
		mockRoom.Duration = 10
		mockRoom.User_id = 1
		mockRoom.Status = "OPEN"
		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Insert Room_id 3", func(t *testing.T) {
		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 30000
		mockRoom.Duration = 3
		mockRoom.User_id = 1
		mockRoom.Status = "OPEN"
		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)
	})
	t.Run("Insert Room_id 4", func(t *testing.T) {
		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 30000
		mockRoom.Duration = 3
		mockRoom.User_id = 1
		mockRoom.Status = "OPEN"
		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Insert Transaction Room_id 1 User_id 2", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 1
		invoice := "INV-N/2/book/41"
		_, err := bookingRepo.CreateTransactions(mockBook.User_id, mockBook.Room_id, invoice, 0)
		assert.Nil(t, err)
	})
	t.Run("Insert Book Room_id 1 User_id 2", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 1
		mockBook.Transaction_id = 1
		_, err := bookingRepo.Create(mockBook)
		assert.Nil(t, err)
	})
	t.Run("Update Transaction Room_id 1 User_id 2 settlement", func(t *testing.T) {
		invoice := "INV-N/2/book/41"
		status := "settlement"
		_, err := transactionRepo.Update(invoice, status)
		assert.Nil(t, err)
	})

	t.Run("Insert Transaction Room_id 2 User_id 2", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 2
		invoice := "INV-N/2/book/42"
		_, err := bookingRepo.CreateTransactions(mockBook.User_id, mockBook.Room_id, invoice, 0)
		assert.Nil(t, err)
	})
	t.Run("Insert Book Room_id 2 User_id 2", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 2
		mockBook.Transaction_id = 2
		_, err := bookingRepo.Create(mockBook)
		assert.Nil(t, err)
	})
	t.Run("Update Transaction Room_id 2 User_id 2 not", func(t *testing.T) {
		invoice := "INV-N/2/book/42"
		status := "not"
		_, err := transactionRepo.Update(invoice, status)
		assert.Nil(t, err)
	})
	t.Run("Update Transaction Room_id 2 User_id 2 cancel", func(t *testing.T) {
		invoice := "INV-N/2/book/42"
		status := "cancel"
		_, err := transactionRepo.Update(invoice, status)
		assert.Nil(t, err)
	})

	t.Run("Insert Transaction Room_id 3 User_id 2", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 3
		invoice := "INV-N/2/book/43"
		_, err := bookingRepo.CreateTransactions(mockBook.User_id, mockBook.Room_id, invoice, 0)
		assert.Nil(t, err)
	})
	t.Run("Insert Book Room_id 3 User_id 2", func(t *testing.T) {
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 3
		mockBook.Transaction_id = 3
		_, err := bookingRepo.Create(mockBook)
		assert.Nil(t, err)
	})
	t.Run("Update Transaction Room_id 3 User_id 2 settlement", func(t *testing.T) {
		invoice := "INV-N/2/book/43"
		status := "settlement"
		_, err := transactionRepo.Update(invoice, status)
		assert.Nil(t, err)
	})
	t.Run("Update Transaction Room_id 3 User_id 2 settlement duration 3", func(t *testing.T) {
		invoice := "INV-D/2/3/book/42"
		status := "settlement"
		_, err := transactionRepo.Update(invoice, status)
		assert.Nil(t, err)
	})
	t.Run("Get All PENDING Transaction", func(t *testing.T) {
		_, err := transactionRepo.Get(2)
		assert.Nil(t, err)
	})
	t.Run("Get All Transaction Database", func(t *testing.T) {
		_, err := transactionRepo.Gets(2)
		assert.Nil(t, err)
	})

	//fake untuk day before now
	t.Run("Create Transaction id 4", func(t *testing.T) {
		var fakeTransaction entities.Transaction
		fakeTransaction.ID = 5
		fakeTransaction.Status = "settlement"
		fakeTransaction.Invoice = "INV-N/2/book/44"
		res, err := transactionRepo.Create(fakeTransaction)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Insert Book Room_id 4 User_id 2", func(t *testing.T) {
		var now = time.Now()
		beforenow := now.AddDate(0, 0, -20)
		var mockBook entities.Book
		mockBook.User_id = 2
		mockBook.Room_id = 4
		mockBook.Checkin = now
		mockBook.Checkout = beforenow
		mockBook.Transaction_id = 5
		_, err := bookingRepo.Create(mockBook)
		assert.Nil(t, err)
	})
	t.Run("Update Transaction Room_id 4 User_id 2 settlement duration 3", func(t *testing.T) {
		invoice := "INV-D/2/4/book/44"
		status := "settlement"
		_, err := transactionRepo.Update(invoice, status)
		assert.Nil(t, err)
	})

}
