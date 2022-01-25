package transactions

import "project-airbnb/entities"

type TransactionsInterface interface {
	Get(userID uint) ([]entities.Transaction, error)
	Gets(userID uint) ([]entities.Transaction, error)
	Update(invoiceID, status string) (entities.Transaction, error)
}
