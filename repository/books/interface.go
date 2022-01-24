package books

import "project-airbnb/entities"

type BooksInterface interface {
	Gets(userID uint) ([]entities.Book, error)
	Get(userID, roomID uint) (entities.Book, error)
	Create(newBooking entities.Book) (entities.Book, error)
	CreateTransactions(userID uint, invoiceID string, roomID uint) (string, error)
	Update(trID uint) (entities.Book, error)
}
