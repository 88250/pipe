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
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

const tablePrefix = "b3_pipe_"

var db *gorm.DB

func ConnectDB() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	var err error
	db, err = gorm.Open("sqlite3", util.Conf.DataFilePath)
	if nil != err {
		log.Fatalf("opens database file [%s] failed: "+err.Error(), util.Conf.DataFilePath)
	}

	tables := []interface{}{
		&model.User{}, &model.Article{}, &model.Comment{}, &model.Navigation{}, &model.Tag{},
		&model.Category{}, &model.Setting{}, &model.Correlation{},
	}

	// TODO: D, remove it after release 1.0.0
	if err = db.DropTableIfExists(tables...).Error; nil != err {
		log.Fatal("drops tables failed: " + err.Error())
	}

	if err = db.AutoMigrate(tables...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}

	if err = db.Model(&model.Article{}).AddUniqueIndex("idx_article_path", "path").Error; nil != err {
		log.Fatal("adds index failed: " + err.Error())
	}
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		log.Error("Disconnect from database failed: " + err.Error())
	}
}
