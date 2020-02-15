// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
	"strconv"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
)

func outputAtomAction(c *gin.Context) {
	feed := generateFeed(c)

	feed.WriteAtom(c.Writer)
}

func outputRSSAction(c *gin.Context) {
	feed := generateFeed(c)

	feed.WriteRss(c.Writer)
}

func generateFeed(c *gin.Context) *feeds.Feed {
	blogID := getBlogID(c)

	feedOutputModeSetting := service.Setting.GetSetting(model.SettingCategoryFeed, model.SettingNameFeedOutputMode, blogID)
	blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, blogID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, blogID)
	blogSubtitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogSubtitle, blogID)
	ret := &feeds.Feed{
		Title:       blogTitleSetting.Value,
		Link:        &feeds.Link{Href: blogURLSetting.Value},
		Description: blogSubtitleSetting.Value,
	}

	var items []*feeds.Item
	articles, _ := service.Article.GetArticles("", 1, blogID)
	for _, article := range articles {
		mdResult := util.Markdown(article.Content)
		description := mdResult.AbstractText
		if strconv.Itoa(model.SettingFeedOutputModeValueFull) == feedOutputModeSetting.Value {
			description = mdResult.ContentHTML
		}
		user := service.User.GetUser(article.AuthorID)
		items = append(items, &feeds.Item{
			Title:       article.Title,
			Link:        &feeds.Link{Href: blogURLSetting.Value + article.Path},
			Description: description,
			Author:      &feeds.Author{Name: user.Name},
			Created:     article.CreatedAt,
		})
	}
	ret.Items = items

	return ret
}
