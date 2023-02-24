package tables

import (
	"gorm.io/gorm"
	"time"
)

type TblBooks struct {
	gorm.Model
	Id          uint
	Isbn        string
	Name        string
	Author      string
	Publisher   string
	ReleaseDate time.Time
	Pages       int
	Description string
}
