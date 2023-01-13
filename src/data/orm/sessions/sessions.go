package sessions

import (
	"api/src/data/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenSession() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(orm.DbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
