package books

import (
	"fmt"
	"project-airbnb/entities"
	"time"

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

func (br *BooksRepository) Get(userID, roomID uint) (entities.Book, error) {
	booking := entities.Book{}
	br.db.Where("user_id=? AND room_id=?", userID, roomID).Find(&booking)
	return booking, nil
}

func (br *BooksRepository) Create(newBooking entities.Book) (entities.Book, error) {
	br.db.Save(&newBooking)
	return newBooking, nil
}

func (tr *BooksRepository) CreateTransactions(userID, roomID uint, invoiceID string) (string, error) {

	midtrans.ServerKey = "SB-Mid-server-W-ANVsQXp9S7q65qndszXrcD"
	midtrans.ClientKey = "SB-Mid-client-QVIZg4p30WL2WLy8"
	midtrans.Environment = midtrans.Sandbox

	room := entities.Room{}
	tr.db.Where("id=?", roomID).Find(&room)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  invoiceID,
			GrossAmt: int64(room.Price),
		},
	}
	snapResp, _ := snap.CreateTransaction(req)

	return snapResp.RedirectURL, nil
}

func (br *BooksRepository) Update(trID uint) (entities.Book, error) {
	oldBook := entities.Book{}
	br.db.Where("transaction_id=?", trID).Find(oldBook)

	var now = time.Now()

	oldBook.Checkin = fmt.Sprint(now.Year(), "-", now.Month(), "-", now.Day())
	oldBook.Checkout = fmt.Sprint(now.Year(), "-", now.Month(), "-", now.Day()+7)
	return oldBook, nil
}
