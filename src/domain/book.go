package domain

import "time"

type Book struct {
	BookId      int       `json:"book_id"`
	Isbn        string    `json:"isbn" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Publisher   string    `json:"publisher" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required" time_format:"01/02/06"`
	Pages       int       `json:"pages" binding:"required"`
	Description string    `json:"description" binding:"required"`
}
