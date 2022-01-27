package books

import "project-airbnb/entities"

type BooksInterface interface {
	Gets(userID uint) ([]entities.Book, error)
	Get(userID, roomID uint) ([]entities.Book, error)
	Create(newBooking entities.Book) (entities.Book, error)
	CreateTransactions(userID, roomID uint, invoiceID string, duration int) (entities.Transaction, error)
	Update(bookID uint) (entities.Book, error)
}
