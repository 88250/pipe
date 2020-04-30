// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package service

import (
	"database/sql"
	"os"
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

var db *gorm.DB
var useSQLite, useMySQL, usePostgres bool

// ConnectDB connects to the database.
func ConnectDB() {
	var err error
	useSQLite = false
	useMySQL = false
	usePostgres = false
	if "" != model.Conf.SQLite {
		db, err = gorm.Open("sqlite3", model.Conf.SQLite)
		useSQLite = true
	} else if "" != model.Conf.MySQL {
		db, err = gorm.Open("mysql", model.Conf.MySQL)
		useMySQL = true
	} else if "" != model.Conf.Postgres {
		db, err = gorm.Open("postgres", model.Conf.Postgres)
		usePostgres = true
	} else {
		logger.Fatal("please specify database")
	}
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	if useSQLite {
		logger.Debug("used [SQLite] as underlying database")
	} else if useMySQL {
		logger.Debug("used [MySQL] as underlying database")
	} else {
		logger.Debug("used [Postgres] as underlying database")
	}

	if err = db.AutoMigrate(model.Models...).Error; nil != err {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	if err = db.Model(&model.Article{}).AddIndex("idx_b3_pipe_articles_created_at", "created_at").Error; nil != err {
		logger.Fatal("adds index failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.LogMode(model.Conf.ShowSQL)
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}

// DBStat returns database statistics.
func DBStat() sql.DBStats {
	return db.DB().Stats()
}

// Database returns the underlying database name.
func Database() string {
	if useSQLite {
		return "SQLite"
	}

	return "MySQL"
}
