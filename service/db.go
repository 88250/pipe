// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
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
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"  // mysql
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

var db *gorm.DB
var useSQLite bool

// ConnectDB connects to the database.
func ConnectDB() {
	var err error
	useSQLite = false
	if "" != model.Conf.SQLite {
		db, err = gorm.Open("sqlite3", model.Conf.SQLite)
		useSQLite = true
	} else if "" != model.Conf.MySQL {
		db, err = gorm.Open("mysql", model.Conf.MySQL)
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

	if err = db.AutoMigrate(model.Models...).Error; nil != err {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	if err = db.Model(&model.Article{}).AddIndex("idx_b3_pipe_articles_created_at", "created_at").Error; nil != err {
		logger.Fatal("adds index failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.LogMode(model.Conf.ShowSQL)
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}

// Database returns the underlying database name.
func Database() string {
	if useSQLite {
		return "SQLite"
	}

	return "MySQL"
}
