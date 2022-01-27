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

func (tr *BooksRepository) CreateTransactions(userID, roomID uint, invoiceID string, duration int) (entities.Transaction, error) {

	room := entities.Room{}
	tr.db.Where("id=?", roomID).Find(&room)
	fmt.Println("===> STATUS ROOM <===", room.Status)
	fmt.Print("===> INVOICEID <===", invoiceID)
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

		newRoom := entities.Room{
			Status: "CLOSED",
		}
		tr.db.Where("id=?", roomID).Model(&room).Updates(newRoom)

		newPrice := (room.Price / room.Duration) * duration

		req := &snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  invoiceID,
				GrossAmt: int64(newPrice),
			},
		}
		snapResp, _ := snap.CreateTransaction(req)

		newTransaction := entities.Transaction{}
		newTransaction.Invoice = invoiceID
		newTransaction.Status = "Pending"
		newTransaction.Url = snapResp.RedirectURL

		tr.db.Save(&newTransaction)

		return newTransaction, nil
	}

}

func (tr *BooksRepository) Update(userID, roomID uint, duration int) (entities.Book, error) {
	book := entities.Book{}
	tr.db.Where("user_id=? AND room_id=?", userID, roomID).Find(&book)

	// updateRoom := entities.Book{
	// 	Checkout: "",
	// }

	// tr.db.Where("id=?", book.ID).Model(&book).Updates(updateRoom)
	return book, nil

}
