package transactions

import "project-airbnb/entities"

type GetTransactionsResponseFormat struct {
	Message string               `json:"message"`
	Data    entities.Transaction `json:"data"`
}
