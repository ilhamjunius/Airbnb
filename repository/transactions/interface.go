package transactions

import "project-airbnb/entities"

type TransactionsInterface interface {
	Get(trID uint) (entities.Transaction, error)
}
