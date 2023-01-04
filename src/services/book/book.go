package bookService

import (
	"api/src/data/repository"
	"api/src/domain"
)

type BookError struct {
	DuplicateEntry string ""
}

type BookService struct{}

func (BookService) FindBookById(bookId int) (domain.Book, error) {
	repo := repository.BookRepository{}

	bookValue, _ := repo.FindBookById(bookId)

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
