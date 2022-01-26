package transactions

import (
	"fmt"
	"project-airbnb/entities"
	"time"

	"gorm.io/gorm"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepo(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{db: db}
}

func (tr *TransactionsRepository) Gets(userID uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}
	bookings := []entities.Book{}
	tr.db.Where("user_id=?", userID).Find(&bookings)
	tr.db.Joins("JOIN books ON books.transaction_id=transactions.id").Where("user_id=?", userID).Find(&transaction)
	return transaction, nil
}

func (tr *TransactionsRepository) Get(userID uint) ([]entities.Transaction, error) {
	transaction := []entities.Transaction{}
	bookings := []entities.Book{}
	tr.db.Where("user_id=?", userID).Find(&bookings)
	tr.db.Joins("JOIN books ON books.transaction_id=transactions.id").Where("user_id=? AND status='PENDING'", userID).Find(&transaction)
	return transaction, nil
}

func (tr *TransactionsRepository) Update(invoiceID, status string) (entities.Transaction, error) {
	transactionUpdate := entities.Transaction{}
	bookUpdate := entities.Book{}
	roomUpdate := entities.Room{}

	if status != "settlement" {
		tr.db.Where("invoice=?", invoiceID).Find(&transactionUpdate)
		newTransactions := entities.Transaction{
			Status: status,
		}
		tr.db.Where("invoice=?", invoiceID).Model(&transactionUpdate).Updates(newTransactions)
	} else {
		tr.db.Where("invoice=?", invoiceID).Find(&transactionUpdate)
		tr.db.Where("transaction_id=?", transactionUpdate.ID).Find(&bookUpdate)
		tr.db.Where("id=?", bookUpdate.Room_id).Find(&roomUpdate)

		var now = time.Now()
		date := fmt.Sprint(now.Year(), "-", now.Month().String()[0:3], "-", now.Day())
		checkin, _ := time.Parse("2006-Jan-02", date)

		checkout := checkin.AddDate(0, 0, +roomUpdate.Duration)
		fmt.Println(checkout)
		newBook := entities.Book{
			Checkin:  checkin.String(),
			Checkout: checkout.String(),
		}
		tr.db.Where("id=?", bookUpdate.ID).Model(&bookUpdate).Updates(newBook)
		fmt.Println(newBook)
		newRoom := entities.Room{
			Status: "CLOSED",
		}
		tr.db.Where("id=?", bookUpdate.Room_id).Model(&roomUpdate).Updates(newRoom)

		transactionUpdate.Status = status
		tr.db.Save(&transactionUpdate)
	}

	return transactionUpdate, nil
}
