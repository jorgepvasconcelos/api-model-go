package sessions

import (
	"api/src/data/db_orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenSession() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(db_orm.DbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
