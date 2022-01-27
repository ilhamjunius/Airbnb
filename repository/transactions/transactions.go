package transactions

import (
	"fmt"
	"project-airbnb/entities"
	"strconv"
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
	fmt.Println("===>All Transaction", transaction)
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

	fmt.Println("===> INVOICE FROM CALLBACK", invoiceID)

	if status != "settlement" {
		tr.db.Where("invoice=?", invoiceID).Find(&transactionUpdate)
		newTransactions := entities.Transaction{
			Status: status,
		}
		tr.db.Where("invoice=?", invoiceID).Model(&transactionUpdate).Updates(newTransactions)
		if status == "cancel" {
			newRoom := entities.Room{
				Status: "OPEN",
			}
			tr.db.Where("id=?", bookUpdate.Room_id).Model(&roomUpdate).Updates(newRoom)
		}
	} else {
		if invoiceID[0:5] != "INV-N" {
			tr.db.Where("invoice=?", invoiceID).Find(&transactionUpdate)
			tr.db.Where("room_id=?", invoiceID[8:9]).Find(&bookUpdate)

			var now = time.Now()
			var duration int
			duration, _ = strconv.Atoi(invoiceID[10:11])
			newdat := bookUpdate.Checkout.AddDate(0, 0, +duration)

			if now.Before(bookUpdate.Checkout) {
				updateBook := entities.Book{
					Checkout: newdat,
				}
				tr.db.Where("id=?", bookUpdate.ID).Model(&bookUpdate).Updates(updateBook)
				transactionUpdate.Status = status
				tr.db.Save(&transactionUpdate)
			} else {
				checkout := now.AddDate(0, 0, +duration)
				newBook := entities.Book{
					Checkin:  now,
					Checkout: checkout,
				}
				tr.db.Where("id=?", bookUpdate.ID).Model(&bookUpdate).Updates(newBook)
				transactionUpdate.Status = status
				tr.db.Save(&transactionUpdate)
			}

		} else {
			tr.db.Where("invoice=?", invoiceID).Find(&transactionUpdate)
			tr.db.Where("transaction_id=?", transactionUpdate.ID).Find(&bookUpdate)
			tr.db.Where("id=?", bookUpdate.Room_id).Find(&roomUpdate)

			now := time.Now()
			date := fmt.Sprint(now.Year(), "-", now.Month().String()[0:3], "-", now.Day())
			checkin, _ := time.Parse("2006-Jan-02", date)

			checkout := checkin.AddDate(0, 0, +roomUpdate.Duration)
			newBook := entities.Book{
				Checkin:  checkin,
				Checkout: checkout,
			}
			tr.db.Where("id=?", bookUpdate.ID).Model(&bookUpdate).Updates(newBook)

			transactionUpdate.Status = status
			tr.db.Save(&transactionUpdate)
		}
	}

	return transactionUpdate, nil
}
