// Solo.go - A small and beautiful blogging platform written in golang.
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
	"errors"
	"sync"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/theme"
	"github.com/b3log/solo.go/util"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Init = &initService{
	mutex: &sync.Mutex{},
}

type initService struct {
	mutex  *sync.Mutex
	inited bool
}

type PlatformStatus struct {
	Version string `json:"version"`
	Locale  string `json:"locale"`
	Inited  bool   `json:"inited"`
}

func (srv *initService) Inited() (platformStatus *PlatformStatus, err error) {
	platformStatus = &PlatformStatus{
		Version: util.Version,
		Locale:  "zh_CN",
	}

	localeSetting := &model.Setting{}
	if err = db.Where("name = ? AND value IS NOT NULL AND blog_id = ?", "locale", uint(1)).
		Find(localeSetting).Error; nil != err {
		if gorm.ErrRecordNotFound == err {
			err = nil
			return
		}

		msg := "checks platform init status failed: " + err.Error()
		log.Error(msg)

		return platformStatus, errors.New(msg)
	}

	srv.inited, platformStatus.Inited = true, true
	platformStatus.Locale = localeSetting.Value

	return
}

func (srv *initService) InitPlatform(platformAdmin *model.User) error {
	if srv.inited {
		return nil
	}

	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	blogID := uint(1)

	saCount := 0
	db.Model(&model.User{}).Where(&model.User{BlogID: blogID}).Count(&saCount)
	if 0 < saCount {
		srv.inited = true

		return nil
	}

	log.Debug("Initializing platform")

	tx := db.Begin()

	if err := initPlatformAdmin(tx, platformAdmin, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initSystemSettings(tx, platformAdmin, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initThemeSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initBasicSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initPreferenceSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initSignSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initI18nSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initFeedSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initStatisticSettings(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := initNavigation(tx, blogID); nil != err {
		tx.Rollback()

		return err
	}
	if err := helloWorld(tx, platformAdmin, blogID); nil != err {
		tx.Rollback()

		return err
	}

	tx.Commit()
	log.Debug("Initialized platform")

	return nil
}

func initPlatformAdmin(tx *gorm.DB, admin *model.User, blogID uint) error {
	admin.Role = model.UserRolePlatformAdmin
	admin.ArticleCount, admin.PublishedArticleCount = 1, 1 // article "Hello, World!"
	admin.BlogID = blogID
	admin.Locale = "zh_CN"
	if err := tx.Create(admin).Error; nil != err {
		return err
	}

	blogUser := &model.Correlation{
		ID1:    blogID,
		ID2:    admin.ID,
		Type:   model.CorrelationBlogUser,
		BlogID: blogID,
	}
	if err := tx.Create(blogUser).Error; nil != err {
		return err
	}

	return nil
}

func initNavigation(tx *gorm.DB, blogID uint) error {
	navigation := &model.Navigation{
		Title:      "黑客派",
		URL:        "https://hacpai.com",
		IconURL:    "",
		OpenMethod: model.NavigationOpenMethodBlank,
		Number:     0,
		BlogID:     blogID,
	}
	if err := tx.Create(navigation).Error; nil != err {
		return err
	}

	return nil
}

func helloWorld(tx *gorm.DB, admin *model.User, blogID uint) error {
	content := `欢迎使用 [Solo.go](https://github.com/b3log/solo.go) 博客系统。这是系统自动生成的演示文章，编辑或者删除它，然后开始你的独立博客之旅！

另外，欢迎你加入[黑客与画家的社区](https://hacpai.com)，你可以使用博客账号直接登录！

----

Solo.go 博客系统是一个开源项目，如果你觉得它很赞，请到[项目首页](https://github.com/b3log/solo.go)给颗星鼓励一下 :heart:`

	article := &model.Article{
		AuthorID:     admin.ID,
		Title:        "世界，你好！",
		Abstract:     content,
		Tags:         "Solo.go",
		Content:      content,
		Path:         "/hello-world",
		Status:       model.ArticleStatusPublished,
		Topped:       false,
		Commentable:  true,
		CommentCount: 1,
		BlogID:       blogID,
	}
	if err := tx.Create(article).Error; nil != err {
		return err
	}

	tag := &model.Tag{
		Title:                 "Solo.go",
		ArticleCount:          1,
		PublishedArticleCount: 1,
		BlogID:                blogID,
	}
	if err := tx.Create(tag).Error; nil != err {
		return err
	}

	articleTagRel := &model.Correlation{
		ID1:    article.ID,
		ID2:    tag.ID,
		Type:   model.CorrelationArticleTag,
		BlogID: blogID,
	}
	if err := tx.Create(articleTagRel).Error; nil != err {
		return err
	}

	comment := &model.Comment{
		ArticleID:       article.ID,
		AuthorName:      "Daniel",
		AuthorAvatarURL: "https://img.hacpai.com/avatar/1353745196354_1500432853138.png?imageView2/1/w/80/h/80/interlace/0/q/100",
		Content:         "写博客需要坚持，相信积累后必然会有收获，我们一起努力加油 :smile:",
		BlogID:          blogID,
	}
	if err := tx.Create(comment).Error; nil != err {
		return err
	}

	return nil
}

func initSystemSettings(tx *gorm.DB, blogAdmin *model.User, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategorySystem,
		Name:     model.SettingNameSystemPath,
		Value:    "/" + blogAdmin.Name,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategorySystem,
		Name:     model.SettingNameSystemVer,
		Value:    "1.0.0",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initThemeSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryTheme,
		Name:     model.SettingNameThemeName,
		Value:    theme.DefaultTheme,
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initBasicSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicBlogSubtitle,
		Value:    "小而美的 golang 博客平台",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicBlogTitle,
		Value:    "Solo.go 示例",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicCommentable,
		Value:    "true",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicFooter,
		Value:    "<!-- 这里可用于放置备案信息等，支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicHeader,
		Value:    "<!-- 这里可用于插入第三方统计等，支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicMetaDescription,
		Value:    "小而美的 golang 博客平台",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicMetaKeywords,
		Value:    "Solo.go,golang,博客,开源",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicNoticeBoard,
		Value:    "<!-- 支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initPreferenceSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListPageSize,
		Value:    "20",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListWindowSize,
		Value:    "20",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListStyle,
		Value:    model.SettingPreferenceArticleListStyleValueTitleAbstract,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceExternalArticleListSize,
		Value:    "7",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostCommentArticleListSize,
		Value:    "7",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostUseTagListSize,
		Value:    "15",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostViewArticleListSize,
		Value:    "7",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRandomArticleListSize,
		Value:    "7",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRecentCommentListSize,
		Value:    "7",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRelevantArticleListSize,
		Value:    "7",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initSignSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategorySign,
		Name:     model.SettingNameArticleSign,
		Value:    "<!-- 支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initI18nSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryI18n,
		Name:     model.SettingNameI18nLocale,
		Value:    "zh_CN",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryI18n,
		Name:     model.SettingNameI18nTimezone,
		Value:    "Asia/Shanghai",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initFeedSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryFeed,
		Name:     model.SettingNameFeedOutputSize,
		Value:    "20",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryFeed,
		Name:     model.SettingNameFeedOutputMode,
		Value:    model.SettingFeedOutputModeValueAbstract,
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initStatisticSettings(tx *gorm.DB, blogID uint) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticArticleCount,
		Value:    "1", // article "Hello, World!"
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticPublishedArticleCount,
		Value:    "1", // article "Hello, World!"
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticCommentCount,
		Value:    "1", // comment of article "Hello, World!"
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticViewCount,
		Value:    "0",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}
