package repository

import (
	"time"
)

type BookDTO struct {
	Id          int
	Isbn        string
	Name        string
	Author      string
	Publisher   string
	ReleaseDate time.Time `time_format:"01/02/06"`
	Pages       int
	Description string
}

type BookRepository struct{}

func (BookRepository) FindBookById(bookId int) (BookDTO, error) {
	date := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	layout := "01/02/06"
	releaseDate, _ := time.Parse(layout, date.Format(layout))

	response := BookDTO{
		Id:          10,
		Isbn:        "ddds",
		Name:        "name",
		Author:      "authott",
		Publisher:   "nit",
		ReleaseDate: releaseDate,
		Pages:       5222,
		Description: "a dest",
	}

	return response, nil
}
