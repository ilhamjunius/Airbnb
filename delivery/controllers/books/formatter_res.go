package books

import "project-airbnb/entities"

type BookingsResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.Book `json:"data"`
}

type TransactionsResponseFormat struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    entities.Transaction `json:"data"`
}
