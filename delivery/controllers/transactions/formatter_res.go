package transactions

import "project-airbnb/entities"

type UpdateTransactionsResponseFormat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    entities.Transaction `json:"data"`
}
