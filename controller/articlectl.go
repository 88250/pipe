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

package controller

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/vinta/pangu"
)

func showArticlesAction(c *gin.Context) {
	page := util.GetPage(c)
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	session := util.GetSession(c)
	articleModels, pagination := service.Article.GetArticles(page, blogID)
	articles := []*ThemeArticle{}
	for _, articleModel := range articleModels {
		themeTags := []*ThemeTag{}
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
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

		author := &ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		mdResult := util.Markdown(articleModel.Content)
		article := &ThemeArticle{
			ID:           articleModel.ID,
			Abstract:     mdResult.AbstractText,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        pangu.SpacingText(articleModel.Title),
			Tags:         themeTags,
			URL:          getBlogURL(c) + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
			ThumbnailURL: mdResult.ThumbURL,
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

	themeTags := []*ThemeTag{}
	tagStrs := strings.Split(article.Tags, ",")
	for _, tagStr := range tagStrs {
		themeTag := &ThemeTag{
			Title: tagStr,
			URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
		}
		themeTags = append(themeTags, themeTag)
	}

	mdResult := util.Markdown(article.Content)
	authorModel := service.User.GetUser(article.AuthorID)
	dataModel["Article"] = &ThemeArticle{
		Author: &ThemeAuthor{
			Name:      authorModel.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		},
		ID:           article.ID,
		CreatedAt:    article.CreatedAt.Format("2006-01-02"),
		Title:        pangu.SpacingText(article.Title),
		Tags:         themeTags,
		URL:          getBlogURL(c) + article.Path,
		Topped:       article.Topped,
		ViewCount:    article.ViewCount,
		CommentCount: article.CommentCount,
		Content:      template.HTML(mdResult.ContentHTML),
		Editable:     session.UID == authorModel.ID,
	}

	page := util.GetPage(c)
	commentModels, pagination := service.Comment.GetArticleComments(article.ID, page, blogID)
	comments := []*ThemeComment{}
	for _, commentModel := range commentModels {
		commentAuthor := service.User.GetUser(commentModel.AuthorID)
		if nil == commentAuthor {
			logger.Errorf("not found comment author [userID=%d]", commentModel.AuthorID)

			continue
		}

		commentAuthorBlog := service.User.GetOwnBlog(commentAuthor.ID)
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID)
		commentAuthorURL := blogURLSetting.Value + util.PathAuthors + "/" + commentAuthor.Name
		author := &ThemeAuthor{
			Name:      commentAuthor.Name,
			URL:       commentAuthorURL,
			AvatarURL: commentAuthor.AvatarURL,
		}

		mdResult := util.Markdown(commentModel.Content)
		comment := &ThemeComment{
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
			parentCommentAuthorModel := service.User.GetUser(parentCommentModel.AuthorID)
			page := service.Comment.GetCommentPage(commentModel.ArticleID, commentModel.ID, commentModel.BlogID)

			parentComment := &ThemeComment{
				ID:  parentCommentModel.ID,
				URL: getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(parentCommentModel.ID)),
				Author: &ThemeAuthor{
					Name: parentCommentAuthorModel.Name,
				},
			}
			comment.Parent = parentComment
		}

		comments = append(comments, comment)
	}

	dataModel["Comments"] = comments
	dataModel["Pagination"] = pagination
	dataModel["RecommendArticles"] = recommendArticles()
	fillPreviousArticle(c, article, &dataModel)
	fillNextArticle(c, article, &dataModel)
	dataModel["ToC"] = template.HTML(toc(dataModel["Article"].(*ThemeArticle)))
	c.HTML(http.StatusOK, getTheme(c)+"/article.html", dataModel)

	service.Article.IncArticleViewCount(article)
}

func fillPreviousArticle(c *gin.Context, article *model.Article, dataModel *DataModel) {
	previous := service.Article.GetPreviousArticle(article.ID, article.BlogID)
	if nil == previous {
		return
	}

	author := service.User.GetUser(previous.AuthorID)
	previousArticle := &ThemeArticle{
		Title: previous.Title,
		URL:   getBlogURL(c) + previous.Path,
		Author: &ThemeAuthor{
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
	nextArticle := &ThemeArticle{
		Title: next.Title,
		URL:   getBlogURL(c) + next.Path,
		Author: &ThemeAuthor{
			Name:      author.Name,
			URL:       getBlogURL(c) + util.PathAuthors + "/" + author.Name,
			AvatarURL: author.AvatarURL,
		},
	}
	(*dataModel)["NextArticle"] = nextArticle
}

func toc(article *ThemeArticle) string {
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
		builder.WriteString("<li class='toc-")
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

func recommendArticles() []*ThemeArticle {
	ret := []*ThemeArticle{}

	result := util.NewResult()
	_, _, errs := gorequest.New().Get(util.HacPaiURL+"/apis/recommend/articles").
		Set("user-agent", util.UserAgent).Timeout(30 * time.Second).EndStruct(result)
	if nil != errs {
		logger.Errorf("get recommend articles: %s", errs)

		return ret
	}

	if 0 != result.Code {
		return ret
	}

	for _, entry := range result.Data.([]interface{}) {
		article := entry.(map[string]interface{})
		author := &ThemeAuthor{
			Name:      article["articleAuthorName"].(string),
			URL:       "https://hacpai.com/member/" + article["articleAuthorName"].(string),
			AvatarURL: "",
		}

		ret = append(ret, &ThemeArticle{
			Author:       author,
			CreatedAt:    time.Now().Format("2006-01-02"),
			Title:        article["articleTitle"].(string),
			URL:          article["articlePermalink"].(string),
			CommentCount: int(article["articleCommentCount"].(float64)),
			ThumbnailURL: "https://img.hacpai.com/bing/20171109.jpg?imageView2/1/w/960/h/520/interlace/1/q/100",
			Abstract:     "丁亮快来修改我呀",
		})
	}

	return ret
}
