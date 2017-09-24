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
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Init = &initService{}

type initService struct {
}

func (srv *initService) InitPlatform(sa *model.User) {
	blogID := uint(1)

	log.Debug("Initializing platform")
	tx := db.Begin()

	if nil != initPreference(tx, blogID) {
		tx.Rollback()
	}
	if nil != initAdmin(tx, sa, blogID) {
		tx.Rollback()
	}
	if nil != helloWorld(tx, sa, blogID) {
		tx.Rollback()
	}

	tx.Commit()
	log.Debugf("Initialized blog [id=%s]", blogID)
}

func initAdmin(tx *gorm.DB, admin *model.User, blogID uint) error {
	if 1 == blogID {
		admin.Role = model.UserRolePlatformAdmin
	} else {
		admin.Role = model.UserRoleBlogAdmin
	}

	admin.BlogID = blogID

	if err := tx.Create(admin).Error; nil != err {
		return err
	}

	return nil
}

func helloWorld(tx *gorm.DB, admin *model.User, blogID uint) error {
	content := `欢迎使用 [Solo.go](https://github.com/b3log/solo.go) 博客系统。这是系统自动生成的演示文章，编辑或者删除它，然后开始你的独立博客之旅！\n\
\n\
另外，欢迎你加入[黑客与画家的社区](https://hacpai.com)，你可以使用博客账号直接登录！\n\
\n\
----\n\
\n\
Solo.go 博客系统是一个开源项目，如果你觉得它很赞，请到[项目首页](https://github.com/b3log/solo.go)给颗星鼓励一下 :heart:`

	article := &model.Article{
		AuthorID:    admin.ID,
		Title:       "世界，你好！",
		Abstract:    content,
		Tags:        "Solo.go",
		Content:     content,
		Permalink:   "/hello-world",
		Status:      model.ArticleStatusPublished,
		Topped:      false,
		Commentable: true,
		Password:    "",
		ViewCount:   0,
		BlogID:      blogID,
	}
	if err := tx.Create(article).Error; nil != err {
		return err
	}

	return nil
}

func initPreference(tx *gorm.DB, blogID uint) error {
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
		Name:     model.SettingNamePreferenceBlogSubtitle,
		Value:    "golang 开源博客",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceBlogTitle,
		Value:    "Solo.go 示例",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceCommentable,
		Value:    "true",
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
		Name:     model.SettingNamePreferenceFeedOutputSize,
		Value:    "20",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceFeedOutputMode,
		Value:    model.SettingPreferenceFeedOutputModeValueAbstract,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceFooter,
		Value:    "<!-- 这里可用于放置备案信息等，支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceHeader,
		Value:    "<!-- 这里可用于插入第三方统计等，支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceLocale,
		Value:    "zh_CN",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMetaDes,
		Value:    "小而美的 golang 博客系统",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMetaKey,
		Value:    "Solo.go,golang,博客,开源",
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
		Name:     model.SettingNamePreferenceNoticeBoard,
		Value:    "<!-- 支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuAK,
		Value:    "",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuBucket,
		Value:    "",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuDomain,
		Value:    "",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuSK,
		Value:    "",
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
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceSign,
		Value:    "<!-- 支持 HTML、脚本 -->",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceSkin,
		Value:    "classic",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceTimezone,
		Value:    "Asia/Shanghai",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceVer,
		Value:    "1.0.0",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}
