// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
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

	Category string `sql:"index" gorm:"size:32" json:"category"`
	Name     string `sql:"index" gorm:"size:64" json:"name"`
	Value    string `gorm:"type:text" json:"value"`

	BlogID uint64 `sql:"index" json:"blogID"`
}

// Setting names of category "system".
const (
	SettingCategorySystem = "system"

	SettingNameSystemVer = "systemVersion"
)

// Setting names of category "theme".
const (
	SettingCategoryTheme = "theme"

	SettingNameThemeName = "themeName"
)

// Setting names of category "basic".
const (
	SettingCategoryBasic = "basic"

	SettingNameBasicBlogURL         = "basicBlogURL"
	SettingNameBasicBlogSubtitle    = "basicBlogSubtitle"
	SettingNameBasicBlogTitle       = "basicBlogTitle"
	SettingNameBasicCommentable     = "basicCommentable"
	SettingNameBasicFooter          = "basicFooter"
	SettingNameBasicHeader          = "basicHeader" // Removed from UI since v1.1.0 caused by issue 54 (https://github.com/b3log/pipe/issues/54)
	SettingNameBasicNoticeBoard     = "basicNoticeBoard"
	SettingNameBasicMetaDescription = "basicMetaDescription"
	SettingNameBasicMetaKeywords    = "basicMetaKeywords"
	SettingNameBasicFaviconURL      = "basicFaviconURL"
	SettingNameBasicLogoURL         = "basicLogoURL"
)

// Setting values of category "basic".
const (
	SettingBasicFooterDefault           = "<!-- 这里可用于放置备案信息等，支持 Markdown -->"
	SettingBasicHeaderDefault           = "<!-- https://github.com/b3log/pipe/issues/54 -->"
	SettingBasicBasicNoticeBoardDefault = "<!-- 支持 Markdown -->本博客由 [Pipe](https://github.com/88250/pipe) 强力驱动"
)

// Setting names of category "preference".
const (
	SettingCategoryPreference = "preference"

	SettingNamePreferenceArticleListPageSize        = "preferenceArticleListPageSize"
	SettingNamePreferenceArticleListWindowSize      = "preferenceArticleListWindowSize"
	SettingNamePreferenceArticleListStyle           = "preferenceArticleListStyle"
	SettingNamePreferenceMostCommentArticleListSize = "preferenceMostCommentArticleListSize"
	SettingNamePreferenceMostUseTagListSize         = "preferenceMostUseTagListSize"
	SettingNamePreferenceMostViewArticleListSize    = "preferenceMostViewArticleListSize"
	SettingNamePreferenceRecentCommentListSize      = "preferenceRecentCommentListSize"
	SettingNamePreferenceRecommendArticleListSize   = "preferenceRecommendArticleListSize"
)

// Setting values of category "preference".
const (
	SettingPreferenceArticleListStyleValueTitle         = 0
	SettingPreferenceArticleListStyleValueTitleAbstract = 1
	SettingPreferenceArticleListStyleValueTitleContent  = 2

	SettingPreferenceArticleListPageSizeDefault        = 20
	SettingPreferenceArticleListWindowSizeDefault      = 7
	SettingPreferenceArticleListStyleDefault           = SettingPreferenceArticleListStyleValueTitleAbstract
	SettingPreferenceMostCommentArticleListSizeDefault = 7
	SettingPreferenceMostUseTagListSizeDefault         = 15
	SettingPreferenceMostViewArticleListSizeDefault    = 15
	SettingPreferenceRecentCommentListSizeDefault      = 7
	SettingPreferenceRecommendArticleListSizeDefault   = 1
)

// Setting names of category "sign".
const (
	SettingCategorySign = "sign"

	SettingNameArticleSign = "signArticle"
)

// Setting values of category "sign".
const (
	SettingArticleSignDefault = "<!-- 支持 Markdown；可用变量 {title}, {author}, {url} -->"
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

	SettingNameFeedOutputMode = "feedOutputMode"
)

// Setting values of category "feed".
const (
	SettingFeedOutputModeValueAbstract = 0
	SettingFeedOutputModeValueFull     = 1
)

// Setting names of category "thirdStatistic".
const (
	SettingCategoryThirdStatistic = "thirdStatistic"

	SettingNameThirdStatisticBaidu = "thirdStatisticBaidu"
)

// Setting names of category "statistic".
const (
	SettingCategoryStatistic = "statistic"

	SettingNameStatisticArticleCount = "statisticArticleCount"
	SettingNameStatisticCommentCount = "statisticCommentCount"
	SettingNameStatisticViewCount    = "statisticViewCount"
)

// Setting names of category "ad".
const (
	SettingCategoryAd = "ad"

	SettingNameAdGoogleAdSenseArticleEmbed = "adGoogleAdSenseArticleEmbed"
)
