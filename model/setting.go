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

package model

import "github.com/jinzhu/gorm"

// Setting names of category "preference".
const (
	SettingCategoryPreference = "preference"

	SettingNamePreferenceArticleListPageSize        = "articleListPageSize"
	SettingNamePreferenceArticleListWindowSize      = "articleListWindowSize"
	SettingNamePreferenceArticleListStyle           = "articleListStyle"
	SettingNamePreferenceBlogSubtitle               = "blogSubtitle"
	SettingNamePreferenceBlogTitle                  = "blogTitle"
	SettingNamePreferenceCommentable                = "commentable"
	SettingNamePreferenceExternalArticleListSize    = "externalRelevantArticleListSize"
	SettingNamePreferenceFeedOutputSize             = "feedOutputSize"
	SettingNamePreferenceFeedOutputMode             = "feedOutputMode"
	SettingNamePreferenceFooter                     = "footer"
	SettingNamePreferenceHeader                     = "header"
	SettingNamePreferenceLocale                     = "locale"
	SettingNamePreferenceMetaDes                    = "metaDes"
	SettingNamePreferenceMetaKey                    = "metaKey"
	SettingNamePreferenceMostCommentArticleListSize = "mostCommentArticleListSize"
	SettingNamePreferenceMostUseTagListSize         = "mostUseTagListSize"
	SettingNamePreferenceMostViewArticleListSize    = "mostViewArticleListSize"
	SettingNamePreferenceNoticeBoard                = "noticeBoard"
	SettingNamePreferenceQiniuAK                    = "qiniuAK"
	SettingNamePreferenceQiniuBucket                = "qiniuBucket"
	SettingNamePreferenceQiniuDomain                = "qiniuDomain"
	SettingNamePreferenceQiniuSK                    = "qiniuSK"
	SettingNamePreferenceRandomArticleListSize      = "randomArticleListSize"
	SettingNamePreferenceRecentCommentListSize      = "recentCommentListSize"
	SettingNamePreferenceRelevantArticleListSize    = "relevantArticleListSize"
	SettingNamePreferenceSign                       = "sign"
	SettingNamePreferenceSkin                       = "skin"
	SettingNamePreferenceTimezone                   = "timezone"
	SettingNamePreferenceVer                        = "version"
)

// Setting model.
type Setting struct {
	gorm.Model

	Category string `gorm:"size:32"`
	Name     string `gorm:"size:32"`
	Value    string `gorm:"type:text"`

	TenantID uint
}
