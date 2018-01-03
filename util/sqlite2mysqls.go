package util

import (
	"github.com/b3log/pipe/model"
	"github.com/jinzhu/gorm"
)

func sqlite2MySQL(sqliteDataFilePath, mysqlConn string) error {
	sqlite, err := gorm.Open("sqlite3", Conf.SQLite)
	if nil != err {
		logger.Fatalf("opens SQLite database failed: " + err.Error())
	}
	mysql, err := gorm.Open("mysql", Conf.MySQL)
	if nil != err {
		logger.Fatalf("opens MySQL database failed: " + err.Error())
	}

	rows := []*model.User{}
	if err = sqlite.Find(&rows).Error; nil != err {
		logger.Errorf("query data failed: ", err.Error())
	}

	for _, row := range rows {
		logger.Info(row)
	}

	_ = mysql

	return nil
}
