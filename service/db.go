// Pipe - A small and beautiful blogging platform written in golang.
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
	"os"

	"github.com/b3log/pipe/log"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Logger
var logger = log.NewLogger(os.Stdout)

const tablePrefix = "b3_pipe_"

var db *gorm.DB

func ConnectDB() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	var err error
	db, err = gorm.Open("sqlite3", util.Conf.DataFilePath)
	if nil != err {
		logger.Fatalf("opens database file [%s] failed: "+err.Error(), util.Conf.DataFilePath)
	}

	// TODO: D, remove it after release 1.0.0
	tables := []interface{}{
		&model.User{}, &model.Article{}, &model.Comment{}, &model.Navigation{}, &model.Tag{},
		&model.Category{}, &model.Archive{}, &model.Setting{}, &model.Correlation{},
	}

	//	if err = db.DropTableIfExists(tables...).Error; nil != err {
	//		logger.Fatal("drops tables failed: " + err.Error())
	//	}

	if err = db.AutoMigrate(tables...).Error; nil != err {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	if err = db.Model(&model.Article{}).AddIndex("idx_b3_pipe_articles_created_at", "created_at").Error; nil != err {
		logger.Fatal("adds index failed: " + err.Error())
	}

	db.LogMode(true)
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}
