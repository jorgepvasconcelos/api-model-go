package bookService

import (
	"api/src/data/errors/sqlError"
	"api/src/data/repository/bookRepository"
	"api/src/domain"
	"api/src/services/book/bookError"
	"errors"
)

func FindBookById(bookId int) (domain.Book, bookError.BookError) {
	bookValue, err := bookRepository.FindBookById(bookId)
	if err != nil {
		if errors.Is(err, sqlError.NotFound) {
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
