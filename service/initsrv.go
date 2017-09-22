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
)

var Init = &initService{}

type initService struct {
}

func (srv *initService) InitBlog() {
	tx := db.Begin()

	if nil != initPreference(tx) {
		tx.Rollback()
	}

	tx.Commit()
}

func initPreference(tx *gorm.DB) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListPageSize,
		Value:    "20"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListWindowSize,
		Value:    "20"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListStyle,
		Value:    model.SettingPreferenceArticleListStyleValueTitleAbstract}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceBlogSubtitle,
		Value:    "golang 开源博客"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceBlogTitle,
		Value:    "Solo.go 示例"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceCommentable,
		Value:    "true"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceExternalArticleListSize,
		Value:    "7"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceFeedOutputSize,
		Value:    "20"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceFeedOutputMode,
		Value:    model.SettingPreferenceFeedOutputModeValueAbstract}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceFooter,
		Value:    "<!-- 这里可用于放置备案信息等，支持脚本 -->"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceHeader,
		Value:    "<!-- 这里可用于插入第三方统计等，支持脚本 -->"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceLocale,
		Value:    "zh_CN"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMetaDes,
		Value:    "小而美的 golang 博客系统"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMetaKey,
		Value:    "Solo.go,golang,博客,开源"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostCommentArticleListSize,
		Value:    "7"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostUseTagListSize,
		Value:    "15"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostViewArticleListSize,
		Value:    "7"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceNoticeBoard,
		Value:    "<!-- 公告栏，支持脚本 -->"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuAK,
		Value:    ""}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuBucket,
		Value:    ""}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuDomain,
		Value:    ""}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceQiniuSK,
		Value:    ""}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRandomArticleListSize,
		Value:    "7"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRecentCommentListSize,
		Value:    "7"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRelevantArticleListSize,
		Value:    "7"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceSign,
		Value:    "默认的签名档"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceSkin,
		Value:    "classic"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceTimezone,
		Value:    "Asia/Shanghai"}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceVer,
		Value:    "1.0.0"}).Error; nil != err {
		return err
	}

	return nil
}
