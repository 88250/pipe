package util

import (
	"github.com/b3log/pipe/model"
	"github.com/jinzhu/gorm"
)

func sqlite2MySQL(sqliteDataFilePath, mysqlConn string) {
	sqlite, err := gorm.Open("sqlite3", Conf.SQLite)
	if nil != err {
		logger.Fatalf("opens SQLite database failed: " + err.Error())
	}
	mysql, err := gorm.Open("mysql", Conf.MySQL)
	if nil != err {
		logger.Fatalf("opens MySQL database failed: " + err.Error())
	}
	if err = mysql.AutoMigrate(Models...).Error; nil != err {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	importArchives(sqlite, mysql, []*model.Archive{})
	importArticles(sqlite, mysql, []*model.Article{})
	importCategories(sqlite, mysql, []*model.Category{})
	importComments(sqlite, mysql, []*model.Comment{})
	importCorrelations(sqlite, mysql, []*model.Correlation{})
	importNavigations(sqlite, mysql, []*model.Navigation{})
	importSettings(sqlite, mysql, []*model.Setting{})
	importTags(sqlite, mysql, []*model.Tag{})
	importUsers(sqlite, mysql, []*model.User{})
}

func importArchives(sqlite, mysql *gorm.DB, models []*model.Archive) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] archives", len(models))
}

func importArticles(sqlite, mysql *gorm.DB, models []*model.Article) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] articles", len(models))
}

func importCategories(sqlite, mysql *gorm.DB, models []*model.Category) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] categories", len(models))
}

func importComments(sqlite, mysql *gorm.DB, models []*model.Comment) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] comments", len(models))
}

func importCorrelations(sqlite, mysql *gorm.DB, models []*model.Correlation) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] correlations", len(models))
}

func importNavigations(sqlite, mysql *gorm.DB, models []*model.Navigation) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] navigations", len(models))
}

func importSettings(sqlite, mysql *gorm.DB, models []*model.Setting) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] settings", len(models))
}

func importTags(sqlite, mysql *gorm.DB, models []*model.Tag) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] tags", len(models))
}

func importUsers(sqlite, mysql *gorm.DB, models []*model.User) {
	if err := sqlite.Find(&models).Error; nil != err {
		logger.Fatalf("queries data failed: ", err.Error())
	}
	for _, model := range models {
		if err := mysql.Save(model).Error; nil != err {
			logger.Fatalf("saves data failed: ", err.Error())
		}
	}
	logger.Infof("imported [%d] users", len(models))
}
