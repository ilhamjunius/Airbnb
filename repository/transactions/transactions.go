package transactions

import (
	"project-airbnb/entities"

	"gorm.io/gorm"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepo(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{db: db}
}

func (tr *TransactionsRepository) Get(trID uint) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	tr.db.Where("id=?", trID).Find(&transaction)
	return transaction, nil
}
