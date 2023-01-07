package dto

import "time"

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
