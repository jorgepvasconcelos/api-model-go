package bookRepository

import (
	"api/src/data/db_orm/sessions"
	"api/src/data/db_orm/tables"
	"api/src/data/dto"
	"api/src/data/errors/sqlError"
	"errors"
	"gorm.io/gorm"
)

func FindBookById(bookId int) (dto.BookDTO, sqlError.SqlError) {
	session, _ := sessions.OpenSession()

	var tblBook tables.TblBooks
	result := session.First(&tblBook, "id = ?", rune(bookId))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dto.BookDTO{}, sqlError.NotFound
	}

	response := dto.BookDTO{
		Id:          int(tblBook.ID),
		Isbn:        tblBook.Isbn,
		Name:        tblBook.Name,
		Author:      tblBook.Author,
		Publisher:   tblBook.Publisher,
		ReleaseDate: tblBook.ReleaseDate,
		Pages:       tblBook.Pages,
		Description: tblBook.Description,
	}

	return response, nil
}
