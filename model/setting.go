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

package model

// Setting model.
type Setting struct {
	Model

	Category string `gorm:"size:32" json:"category"`
	Name     string `gorm:"size:32" json:"name"`
	Value    string `gorm:"type:text" json:"value"`

	BlogID uint `json:"blogID"`
}

// Setting names of category "system".
const (
	SettingCategorySystem = "system"

	SettingNameSystemPath = "path"
	SettingNameSystemVer  = "version"
)

// Setting names of category "theme".
const (
	SettingCategoryTheme = "theme"

	SettingNameThemeName = "name"
)

// Setting names of catgory "basic".
const (
	SettingCategoryBasic = "basic"

	SettingNameBasicBlogSubtitle    = "blogSubtitle"
	SettingNameBasicBlogTitle       = "blogTitle"
	SettingNameBasicCommentable     = "commentable"
	SettingNameBasicFooter          = "footer"
	SettingNameBasicHeader          = "header"
	SettingNameBasicNoticeBoard     = "noticeBoard"
	SettingNameBasicMetaDescription = "metaDescription"
	SettingNameBasicMetaKeywords    = "metaKeywords"
)

// Setting names of category "preference".
const (
	SettingCategoryPreference = "preference"

	SettingNamePreferenceArticleListPageSize        = "articleListPageSize"
	SettingNamePreferenceArticleListWindowSize      = "articleListWindowSize"
	SettingNamePreferenceArticleListStyle           = "articleListStyle"
	SettingNamePreferenceExternalArticleListSize    = "externalRelevantArticleListSize"
	SettingNamePreferenceMostCommentArticleListSize = "mostCommentArticleListSize"
	SettingNamePreferenceMostUseTagListSize         = "mostUseTagListSize"
	SettingNamePreferenceMostViewArticleListSize    = "mostViewArticleListSize"
	SettingNamePreferenceRandomArticleListSize      = "randomArticleListSize"
	SettingNamePreferenceRecentCommentListSize      = "recentCommentListSize"
	SettingNamePreferenceRelevantArticleListSize    = "relevantArticleListSize"
	SettingNamePreferenceTheme                      = "theme"
	SettingNamePreferenceVer                        = "version"
)

// Setting names of category "sign".
const (
	SettingCategorySign = "sign"

	SettingNameArticleSign = "articleSign"
)

// Setting names of category "i18n".
const (
	SettingCategoryI18n = "i18n"

	SettingNameI18nLocale   = "locale"
	SettingNameI18nTimezone = "timezone"
)

// Setting names of category "feed".
const (
	SettingCategoryFeed = "feed"

	SettingNameFeedOutputSize = "outputSize"
	SettingNameFeedOutputMode = "outputMode"
)

// Setting value of category "feed".
const (
	SettingFeedOutputModeValueAbstract = "abstract"
	SettingFeedOutputModeValueFull     = "full"
)

// Setting names of category "statistic".
const (
	SettingCategoryStatistic = "statistic"

	SettingNameStatisticArticleCount          = "articleCount"
	SettingNameStatisticPublishedArticleCount = "publishedArticleCount"
	SettingNameStatisticCommentCount          = "commentCount"
	SettingNameStatisticViewCount             = "viewCount"
)

// Setting values of category "preference".
const (
	SettingPreferenceArticleListStyleValueTitle         = "title"
	SettingPreferenceArticleListStyleValueTitleAbstract = "titleAbstract"
	SettingPreferenceArticleListStyleValueTitleContent  = "titleContent"
)
