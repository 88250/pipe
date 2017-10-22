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

	SettingNameSystemPath = "systemPath"
	SettingNameSystemVer  = "systemVersion"
)

// Setting names of category "theme".
const (
	SettingCategoryTheme = "theme"

	SettingNameThemeName = "themeName"
)

// Setting names of catgory "basic".
const (
	SettingCategoryBasic = "basic"

	SettingNameBasicBlogSubtitle    = "basicBlogSubtitle"
	SettingNameBasicBlogTitle       = "basicBlogTitle"
	SettingNameBasicCommentable     = "basicCommentable"
	SettingNameBasicFooter          = "basicFooter"
	SettingNameBasicHeader          = "basicHeader"
	SettingNameBasicNoticeBoard     = "basicNoticeBoard"
	SettingNameBasicMetaDescription = "basicMetaDescription"
	SettingNameBasicMetaKeywords    = "basicMetaKeywords"
	SettingNameBasicFaviconURL      = "basicFaviconURL"
)

// Setting names of category "preference".
const (
	SettingCategoryPreference = "preference"

	SettingNamePreferenceArticleListPageSize        = "preferenceArticleListPageSize"
	SettingNamePreferenceArticleListWindowSize      = "preferenceArticleListWindowSize"
	SettingNamePreferenceArticleListStyle           = "preferenceArticleListStyle"
	SettingNamePreferenceExternalArticleListSize    = "preferenceExternalRelevantArticleListSize"
	SettingNamePreferenceMostCommentArticleListSize = "preferenceMostCommentArticleListSize"
	SettingNamePreferenceMostUseTagListSize         = "preferenceMostUseTagListSize"
	SettingNamePreferenceMostViewArticleListSize    = "preferenceMostViewArticleListSize"
	SettingNamePreferenceRandomArticleListSize      = "preferenceRandomArticleListSize"
	SettingNamePreferenceRecentCommentListSize      = "preferenceRecentCommentListSize"
	SettingNamePreferenceRelevantArticleListSize    = "preferenceRelevantArticleListSize"
)

// Setting names of category "sign".
const (
	SettingCategorySign = "sign"

	SettingNameArticleSign = "signArticle"
)

// Setting names of category "i18n".
const (
	SettingCategoryI18n = "i18n"

	SettingNameI18nLocale   = "i18nLocale"
	SettingNameI18nTimezone = "i18nTimezone"
)

// Setting names of category "feed".
const (
	SettingCategoryFeed = "feed"

	SettingNameFeedOutputSize = "feedOutputSize"
	SettingNameFeedOutputMode = "feedOutputMode"
)

// Setting value of category "feed".
const (
	SettingFeedOutputModeValueAbstract = "abstract"
	SettingFeedOutputModeValueFull     = "full"
)

// Setting names of category "statistic".
const (
	SettingCategoryStatistic = "statistic"

	SettingNameStatisticArticleCount          = "statisticArticleCount"
	SettingNameStatisticPublishedArticleCount = "statisticPublishedArticleCount"
	SettingNameStatisticCommentCount          = "statisticCommentCount"
	SettingNameStatisticViewCount             = "statisticViewCount"
)

// Setting values of category "preference".
const (
	SettingPreferenceArticleListStyleValueTitle         = "title"
	SettingPreferenceArticleListStyleValueTitleAbstract = "titleAbstract"
	SettingPreferenceArticleListStyleValueTitleContent  = "titleContent"

	SettingPreferenceArticleListPageSizeDefault        = 20
	SettingPreferenceArticleListWindowSizeDefault      = 20
	SettingPreferenceArticleListStyleDefault           = SettingPreferenceArticleListStyleValueTitleAbstract
	SettingPreferenceExternalArticleListSizeDefault    = 7
	SettingPreferenceMostCommentArticleListSizeDefault = 7
	SettingPreferenceMostUseTagListSizeDefault         = 15
	SettingPreferenceMostViewArticleListSizeDefault    = 15
	SettingPreferenceRandomArticleListSizeDefault      = 7
	SettingPreferenceRecentCommentListSizeDefault      = 7
	SettingPreferenceRelevantArticleListSizeDefault    = 7
)
