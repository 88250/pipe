package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return TABLE_PREFIX + defaultTableName
	}

	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/solo.go?charset=utf8&parseTime=true&loc=Local")
	if nil != err {
		fmt.Println(err)
	}

	db.DropTableIfExists(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	db.AutoMigrate(&User{})
}
