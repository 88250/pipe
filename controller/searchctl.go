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
	"html/template"
	"net/http"
	"strings"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/vinta/pangu"
)

func showOpensearchAction(c *gin.Context) {
	blogID := getBlogID(c)
	blogTitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogTitle, blogID)
	blogSubtitleSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogSubtitle, blogID)

	opensearch := `<?xml version="1.0" encoding="UTF-8" ?>
<OpenSearchDescription xmlns="http://a9.com/-/spec/opensearch/1.1/" xmlns:moz="http://www.mozilla.org/2006/browser/search/">
    <ShortName>${title}</ShortName>
    <Description>${description}</Description>
    <InputEncoding>UTF-8</InputEncoding>
    <Image width="16" height="16" type="image/x-icon">${faviconURL}</Image>
    <Url type="text/html" method="get" template="${blogURL}/search?key={searchTerms}"></Url>
</OpenSearchDescription>
`
	opensearch = strings.Replace(opensearch, "${title}", blogTitleSetting.Value, -1)
	opensearch = strings.Replace(opensearch, "${description}", blogSubtitleSetting.Value, -1)
	faviconURL := getDataModel(c)["FaviconURL"].(string)
	opensearch = strings.Replace(opensearch, "${faviconURL}", faviconURL, -1)
	blogURL := getBlogURL(c)
	opensearch = strings.Replace(opensearch, "${blogURL}", blogURL, -1)

	c.Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
	c.Writer.Write([]byte(opensearch))
}

func searchAction(c *gin.Context) {
	blogID := getBlogID(c)
	key := c.Query("key")
	page := util.GetPage(c)
	articleModels, pagination := service.Article.GetArticles(key, page, blogID)
	var articles []*model.ThemeArticle
	for _, articleModel := range articleModels {
		var themeTags []*model.ThemeTag
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &model.ThemeTag{
				Title: tagStr,
				URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
			}
			themeTags = append(themeTags, themeTag)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		if nil == authorModel {
			logger.Errorf("not found author of article [id=%d, authorID=%d]", articleModel.ID, articleModel.AuthorID)

			continue
		}

		mdResult := util.Markdown(articleModel.Content)
		article := &model.ThemeArticle{
			Title:    pangu.SpacingText(articleModel.Title),
			Abstract: template.HTML(mdResult.AbstractText),
			URL:      getBlogURL(c) + articleModel.Path,
			Tags:     themeTags,
		}

		articles = append(articles, article)
	}

	dataModel := getDataModel(c)
	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	dataModel["Key"] = key
	c.HTML(http.StatusOK, "search.html", dataModel)
}
