// database.go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	//db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(Group{}, Transaction{}, Person{}, Verisimilitude{})
	return db
}
