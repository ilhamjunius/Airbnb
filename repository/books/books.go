package books

import (
	"fmt"
	"project-airbnb/entities"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type BooksRepository struct {
	db *gorm.DB
}

func NewBooksRepo(db *gorm.DB) *BooksRepository {
	return &BooksRepository{db: db}
}

func (br *BooksRepository) Gets(userID uint) ([]entities.Book, error) {
	bookings := []entities.Book{}
	br.db.Where("user_id=?", userID).Find(&bookings)
	return bookings, nil
}

func (br *BooksRepository) Get(userID, roomID uint) ([]entities.Book, error) {
	booking := []entities.Book{}
	br.db.Where("user_id=? AND room_id=?", userID, roomID).Find(&booking)

	return booking, nil
}

func (br *BooksRepository) Create(newBooking entities.Book) (entities.Book, error) {
	br.db.Save(&newBooking)
	return newBooking, nil
}

func (tr *BooksRepository) CreateTransactions(userID, roomID uint, invoiceID string) (entities.Transaction, error) {

	room := entities.Room{}
	tr.db.Where("id=?", roomID).Find(&room)
	fmt.Println("===> STATUS ROOM <===", room.Status)
	if room.Status != "CLOSED" {
		newRoom := entities.Room{
			Status: "CLOSED",
		}
		tr.db.Where("id=?", roomID).Model(&room).Updates(newRoom)

		req := &snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  invoiceID,
				GrossAmt: int64(room.Price),
			},
		}
		snapResp, _ := snap.CreateTransaction(req)

		newTransaction := entities.Transaction{}
		newTransaction.Invoice = invoiceID
		newTransaction.Status = "Pending"
		newTransaction.Url = snapResp.RedirectURL

		tr.db.Save(&newTransaction)

		return newTransaction, nil
	} else {
		failTransaction := entities.Transaction{
			ID: 0,
		}
		return failTransaction, nil
	}

}
