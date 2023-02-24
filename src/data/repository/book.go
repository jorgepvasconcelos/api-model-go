package repository

import (
	"api/src/data/db_orm/sessions"
	"api/src/data/db_orm/tables"
	"api/src/data/dto"
)

type BookRepository struct{}

func (BookRepository) FindBookById(bookId int) (dto.BookDTO, error) {

	session, _ := sessions.OpenSession()

	var tblBook tables.TblBooks
	session.First(&tblBook, "id = ?", rune(bookId))

	//date := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	//layout := "01/02/06"
	//releaseDate, _ := time.Parse(layout, date.Format(layout))
	//
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
