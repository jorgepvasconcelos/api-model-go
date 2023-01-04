package tables

import (
	"gorm.io/gorm"
)

type TblBooks struct {
	gorm.Model
	Book  string
	Pages uint
}
