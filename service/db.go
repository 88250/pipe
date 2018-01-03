// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"os"

	"github.com/b3log/pipe/log"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Logger
var logger = log.NewLogger(os.Stdout)

var db *gorm.DB

func ConnectDB() {
	var err error
	useSQLite := false
	if "" != util.Conf.SQLite {
		db, err = gorm.Open("sqlite3", util.Conf.SQLite)
		useSQLite = true
	} else if "" != util.Conf.MySQL {
		db, err = gorm.Open("mysql", util.Conf.MySQL)
	} else {
		logger.Fatal("please specify database")
	}
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	if useSQLite {
		logger.Debug("used [SQLite] as underlying database")
	} else {
		logger.Debug("used [MySQL] as underlying database")
	}

	if err = db.AutoMigrate(util.Models...).Error; nil != err {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	if err = db.Model(&model.Article{}).AddIndex("idx_b3_pipe_articles_created_at", "created_at").Error; nil != err {
		logger.Fatal("adds index failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	//db.LogMode(true)
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}
