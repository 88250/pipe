// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

const tablePrefix = "b3_solo_go_"

var db *gorm.DB

func ConnectDB() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db, err := gorm.Open("sqlite3", util.Conf.DataFilePath)
	if nil != err {
		log.Fatalf("opens database file [%s] failed: "+err.Error(), util.Conf.DataFilePath)
	}

	tables := []interface{}{
		&model.User{}, &model.Article{}, &model.Comment{}, &model.Page{}, &model.Tag{},
		&model.Category{}, &model.Setting{}, &model.Correlation{},
	}

	db.DropTableIfExists(tables...)
	db.AutoMigrate(tables...)
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Error("Disconnect from database failed: " + err.Error())
	}
}
