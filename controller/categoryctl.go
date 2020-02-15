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
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/88250/pipe/i18n"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/vinta/pangu"
)

func showCategoriesAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	locale := getLocale(c)
	categoryModels := service.Category.GetCategories(math.MaxInt8, blogID)
	var themeCategories []*model.ThemeCategory
	for _, categoryModel := range categoryModels {
		var themeTags []*model.ThemeTag
		tagStrs := strings.Split(categoryModel.Tags, ",")
		for _, tagTitle := range tagStrs {
			themeTag := &model.ThemeTag{
				Title: tagTitle,
				URL:   getBlogURL(c) + util.PathTags + "/" + tagTitle,
			}
			themeTags = append(themeTags, themeTag)
		}

		themeCategory := &model.ThemeCategory{
			Title:        categoryModel.Title,
			URL:          getBlogURL(c) + util.PathCategories + categoryModel.Path,
			Description:  categoryModel.Description,
			Tags:         themeTags,
			ArticleCount: service.Category.GetCategoryArticleCount(categoryModel.ID, blogID),
		}
		themeCategories = append(themeCategories, themeCategory)
	}

	dataModel["Categories"] = themeCategories
	dataModel["Title"] = i18n.GetMessage(locale, "categories") + " - " + dataModel["Title"].(string)

	c.HTML(http.StatusOK, getTheme(c)+"/categories.html", dataModel)
}

func showCategoryArticlesArticlesAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	locale := getLocale(c)
	session := util.GetSession(c)
	page := util.GetPage(c)
	categoryPath := strings.SplitAfter(c.Request.URL.Path, util.PathCategories)[1]
	categoryModel := service.Category.GetCategoryByPath(categoryPath, blogID)
	if nil == categoryModel {
		notFound(c)

		return
	}
	articleListStyleSetting := service.Setting.GetSetting(model.SettingCategoryPreference, model.SettingNamePreferenceArticleListStyle, blogID)
	articleModels, pagination := service.Article.GetCategoryArticles(categoryModel.ID, page, blogID)
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

		author := &model.ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		mdResult := util.Markdown(articleModel.Content)
		abstract := template.HTML("")
		thumbnailURL := mdResult.ThumbURL
		if strconv.Itoa(model.SettingPreferenceArticleListStyleValueTitleAbstract) == articleListStyleSetting.Value {
			abstract = template.HTML(mdResult.AbstractText)
		}
		if "" != articleModel.Abstract {
			abstract = template.HTML(articleModel.Abstract)
		}
		if strconv.Itoa(model.SettingPreferenceArticleListStyleValueTitleContent) == articleListStyleSetting.Value {
			abstract = template.HTML(mdResult.ContentHTML)
			thumbnailURL = ""
		}
		article := &model.ThemeArticle{
			ID:             articleModel.ID,
			Abstract:       abstract,
			Author:         author,
			CreatedAt:      articleModel.CreatedAt.Format("2006-01-02"),
			CreatedAtYear:  articleModel.CreatedAt.Format("2006"),
			CreatedAtMonth: articleModel.CreatedAt.Format("01"),
			CreatedAtDay:   articleModel.CreatedAt.Format("02"),
			Title:          pangu.SpacingText(articleModel.Title),
			Tags:           themeTags,
			URL:            getBlogURL(c) + articleModel.Path,
			Topped:         articleModel.Topped,
			ViewCount:      articleModel.ViewCount,
			CommentCount:   articleModel.CommentCount,
			ThumbnailURL:   thumbnailURL,
			Editable:       session.UID == authorModel.ID,
		}

		articles = append(articles, article)
	}
	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	dataModel["Category"] = &model.ThemeCategory{
		Title:        categoryModel.Title,
		ArticleCount: pagination.RecordCount,
	}
	dataModel["Title"] = categoryModel.Title + " - " + i18n.GetMessage(locale, "categories") + " - " + dataModel["Title"].(string)

	c.HTML(http.StatusOK, getTheme(c)+"/category-articles.html", dataModel)
}
