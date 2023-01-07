package bookService

import (
	"api/src/data/repository"
	"api/src/domain"
	"api/src/services/book/bookError"
	"errors"
)

type BookService struct{}

func (BookService) FindBookById(bookId int) (domain.Book, bookError.BookError) {
	bookValue, err := repository.BookRepository{}.FindBookById(bookId)
	if err != nil {
		if errors.Is(err, bookError.NotFound) {
			return domain.Book{}, bookError.NotFound
		}
	}

	book := domain.Book{
		BookId:      bookValue.Id,
		Isbn:        bookValue.Isbn,
		Name:        bookValue.Name,
		Author:      bookValue.Author,
		Publisher:   bookValue.Publisher,
		ReleaseDate: bookValue.ReleaseDate,
		Pages:       bookValue.Pages,
		Description: bookValue.Description,
	}
	return book, nil
}
