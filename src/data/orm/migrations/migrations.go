package main

import (
	"api/src/data/orm"
	"api/src/data/orm/tables"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.Open(orm.DbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrationTables := []interface{}{&tables.TblBooks{}}

	for _, tbl := range migrationTables {
		err = db.Migrator().DropTable(tbl)
		if err != nil {
			log.Fatal(err)
		}
		err = db.Migrator().CreateTable(tbl)
		if err != nil {
			log.Fatal(err)
		}
		//// Migrate the schema
		//err = db.AutoMigrate(&tables.TblBooks{})
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
	fmt.Println("Successs to migrate schema")
}
