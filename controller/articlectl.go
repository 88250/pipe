// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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

package controller

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/b3log/pipe/cron"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/vinta/pangu"
)

func showArticlesAction(c *gin.Context) {
	page := util.GetPage(c)
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	session := util.GetSession(c)
	articleListStyleSetting := service.Setting.GetSetting(model.SettingCategoryPreference, model.SettingNamePreferenceArticleListStyle, blogID)
	articleModels, pagination := service.Article.GetArticles("", page, blogID)
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
		if strconv.Itoa(model.SettingPreferenceArticleListStyleValueTitleContent) == articleListStyleSetting.Value {
			abstract = template.HTML(mdResult.ContentHTML)
			thumbnailURL = ""
		}
		article := &model.ThemeArticle{
			ID:           articleModel.ID,
			Abstract:     abstract,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        pangu.SpacingText(articleModel.Title),
			Tags:         themeTags,
			URL:          getBlogURL(c) + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
			ThumbnailURL: thumbnailURL,
			Editable:     session.UID == authorModel.ID,
		}

		articles = append(articles, article)
	}

	dataModel["Articles"] = articles
	dataModel["Pagination"] = pagination
	c.HTML(http.StatusOK, getTheme(c)+"/index.html", dataModel)
}

func showArticleAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	session := util.GetSession(c)

	a, _ := c.Get("article")
	article := a.(*model.Article)

	var themeTags []*model.ThemeTag
	tagStrs := strings.Split(article.Tags, ",")
	for _, tagStr := range tagStrs {
		themeTag := &model.ThemeTag{
			Title: tagStr,
			URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
		}
		themeTags = append(themeTags, themeTag)
	}

	mdResult := util.Markdown(article.Content)
	authorModel := service.User.GetUser(article.AuthorID)
	articleTitle := pangu.SpacingText(article.Title)
	articleURL := getBlogURL(c) + article.Path
	articleSignSetting := dataModel["Setting"].(map[string]interface{})[model.SettingNameArticleSign].(string)
	articleSignSetting = strings.Replace(articleSignSetting, "{title}", articleTitle, -1)
	articleSignSetting = strings.Replace(articleSignSetting, "{author}", authorModel.Name, -1)
	articleSignSetting = strings.Replace(articleSignSetting, "{url}", articleURL, -1)
	articleSignSetting = util.Markdown(articleSignSetting).ContentHTML
	articleSignSetting = strings.TrimPrefix(articleSignSetting, "<p>")
	articleSignSetting = strings.TrimSuffix(articleSignSetting, "</p>")
	articleSignSetting = strings.TrimSpace(articleSignSetting)
	dataModel["Article"] = &model.ThemeArticle{
		Author: &model.ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		},
		ID:           article.ID,
		Abstract:     template.HTML(mdResult.AbstractText),
		CreatedAt:    article.CreatedAt.Format("2006-01-02"),
		Title:        articleTitle,
		Tags:         themeTags,
		URL:          articleURL,
		Topped:       article.Topped,
		ViewCount:    article.ViewCount,
		CommentCount: article.CommentCount,
		ThumbnailURL: mdResult.ThumbURL,
		Content:      template.HTML(mdResult.ContentHTML + "\n" + articleSignSetting),
		Editable:     session.UID == authorModel.ID,
	}

	page := util.GetPage(c)
	commentModels, pagination := service.Comment.GetArticleComments(article.ID, page, blogID)
	var comments []*model.ThemeComment
	for _, commentModel := range commentModels {
		commentAuthor := service.User.GetUser(commentModel.AuthorID)
		if nil == commentAuthor {
			logger.Errorf("not found comment author [userID=%d]", commentModel.AuthorID)

			continue
		}

		commentAuthorBlog := service.User.GetOwnBlog(commentAuthor.ID)
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID)
		commentAuthorURL := blogURLSetting.Value + util.PathAuthors + "/" + commentAuthor.Name
		author := &model.ThemeAuthor{
			Name:      commentAuthor.Name,
			URL:       commentAuthorURL,
			AvatarURL: commentAuthor.AvatarURL,
		}

		mdResult := util.Markdown(commentModel.Content)
		comment := &model.ThemeComment{
			ID:         commentModel.ID,
			Content:    template.HTML(mdResult.ContentHTML),
			URL:        getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(commentModel.ID)),
			Author:     author,
			CreatedAt:  commentModel.CreatedAt.Format("2006-01-02"),
			Removable:  session.UID == authorModel.ID,
			ReplyCount: service.Comment.GetRepliesCount(commentModel.ID, commentModel.BlogID),
		}
		if 0 != commentModel.ParentCommentID {
			parentCommentModel := service.Comment.GetComment(commentModel.ParentCommentID)
			if nil != parentCommentModel {
				parentCommentAuthorModel := service.User.GetUser(parentCommentModel.AuthorID)
				page := service.Comment.GetCommentPage(commentModel.ArticleID, commentModel.ID, commentModel.BlogID)

				parentComment := &model.ThemeComment{
					ID:  parentCommentModel.ID,
					URL: getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(parentCommentModel.ID)),
					Author: &model.ThemeAuthor{
						Name: parentCommentAuthorModel.Name,
					},
				}
				comment.Parent = parentComment
			}
		}

		comments = append(comments, comment)
	}

	dataModel["Comments"] = comments
	dataModel["Pagination"] = pagination
	dataModel["RecommendArticles"] = getRecommendArticles()
	fillPreviousArticle(c, article, &dataModel)
	fillNextArticle(c, article, &dataModel)
	dataModel["ToC"] = template.HTML(toc(dataModel["Article"].(*model.ThemeArticle)))
	dataModel["Title"] = articleTitle + " - " + dataModel["Title"].(string)

	c.HTML(http.StatusOK, getTheme(c)+"/article.html", dataModel)

	go service.Article.IncArticleViewCount(article)
}

func fillPreviousArticle(c *gin.Context, article *model.Article, dataModel *DataModel) {
	previous := service.Article.GetPreviousArticle(article.ID, article.BlogID)
	if nil == previous {
		return
	}

	author := service.User.GetUser(previous.AuthorID)
	previousArticle := &model.ThemeArticle{
		Title: previous.Title,
		URL:   getBlogURL(c) + previous.Path,
		Author: &model.ThemeAuthor{
			Name:      author.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + author.Name,
			AvatarURL: author.AvatarURL,
		},
	}
	(*dataModel)["PreviousArticle"] = previousArticle
}

func fillNextArticle(c *gin.Context, article *model.Article, dataModel *DataModel) {
	next := service.Article.GetNextArticle(article.ID, article.BlogID)
	if nil == next {
		return
	}

	author := service.User.GetUser(next.AuthorID)
	nextArticle := &model.ThemeArticle{
		Title: next.Title,
		URL:   getBlogURL(c) + next.Path,
		Author: &model.ThemeAuthor{
			Name:      author.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + author.Name,
			AvatarURL: author.AvatarURL,
		},
	}
	(*dataModel)["NextArticle"] = nextArticle
}

func toc(article *model.ThemeArticle) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(article.Content)))
	elements := doc.Find("h1, h2, h3, h4, h5")
	if nil == elements || 3 > elements.Length() {
		return ""
	}

	builder := bytes.Buffer{}
	builder.WriteString("<ul id=\"toc\" class=\"toc\">")
	elements.Each(func(i int, element *goquery.Selection) {
		tagName := goquery.NodeName(element)
		id := "toc_" + tagName + "_" + strconv.Itoa(i)
		element.SetAttr("id", id)
		builder.WriteString("<li class='toc__")
		builder.WriteString(tagName)
		builder.WriteString("'><a href=\"#")
		builder.WriteString(id)
		builder.WriteString("\">")
		builder.WriteString(element.Text())
		builder.WriteString("</a></li>")
	})
	builder.WriteString("</ul>")

	content, _ := doc.Find("body").Html()
	article.Content = template.HTML(content)

	return builder.String()
}

func getRecommendArticles() []*model.ThemeArticle {
	var ret []*model.ThemeArticle

	indics := util.RandInts(0, len(cron.RecommendArticles), 7)
	for _, index := range indics {
		article := cron.RecommendArticles[index]

		ret = append(ret, article)
	}

	return ret
}
