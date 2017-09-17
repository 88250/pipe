package service

import (
	"github.com/b3log/solo.go/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

const tablePrefix = "b3_solo_go_"

var db *gorm.DB

func ConnectDB() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/solo.go?charset=utf8&parseTime=true&loc=Local")
	if nil != err {
		log.Error(err)

		return
	}

	db.DropTableIfExists(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	db.AutoMigrate(&model.User{})
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Error("Disconnect from MySQL failed: " + err.Error())
	}
}
