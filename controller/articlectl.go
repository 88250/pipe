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

package controller

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/88250/gulu"
	"github.com/88250/pipe/cron"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
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
		if "\n" != articleModel.Abstract && "" != articleModel.Abstract {
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
	c.HTML(http.StatusOK, getTheme(c)+"/index.html", dataModel)
}

func showArticleAction(c *gin.Context) {
	dataModel := getDataModel(c)
	blogID := getBlogID(c)
	session := util.GetSession(c)

	a, _ := c.Get("article")
	articleModel := a.(*model.Article)

	var themeTags []*model.ThemeTag
	tagStrs := strings.Split(articleModel.Tags, ",")
	for _, tagStr := range tagStrs {
		themeTag := &model.ThemeTag{
			Title: tagStr,
			URL:   getBlogURL(c) + util.PathTags + "/" + tagStr,
		}
		themeTags = append(themeTags, themeTag)
	}

	mdResult := util.Markdown(articleModel.Content)

	gaSetting := service.Setting.GetSetting(model.SettingCategoryAd, model.SettingNameAdGoogleAdSenseArticleEmbed, blogID)
	if nil != gaSetting && 0 < len(gaSetting.Value) {
		// 嵌入 Google AdSense 文章广告
		if idx := strings.Index(mdResult.ContentHTML, "</p>"); 0 < idx {
			idx = idx + len("</p>")
			gaScript := `
<script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
` + gaSetting.Value + `
<script>
     (adsbygoogle = window.adsbygoogle || []).push({});
</script>
`
			if !strings.Contains(mdResult.ContentHTML, gaScript) {
				mdResult.ContentHTML = mdResult.ContentHTML[0:idx] + gaScript + mdResult.ContentHTML[idx:]
			}
		}
	}

	authorModel := service.User.GetUser(articleModel.AuthorID)
	articleTitle := pangu.SpacingText(articleModel.Title)
	articleURL := getBlogURL(c) + articleModel.Path
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
		ID:             articleModel.ID,
		Abstract:       template.HTML(mdResult.AbstractText),
		CreatedAt:      articleModel.CreatedAt.Format("2006-01-02"),
		CreatedAtYear:  articleModel.CreatedAt.Format("2006"),
		CreatedAtMonth: articleModel.CreatedAt.Format("01"),
		CreatedAtDay:   articleModel.CreatedAt.Format("02"),
		Title:          articleTitle,
		Tags:           themeTags,
		URL:            articleURL,
		Topped:         articleModel.Topped,
		ViewCount:      articleModel.ViewCount,
		CommentCount:   articleModel.CommentCount,
		ThumbnailURL:   mdResult.ThumbURL,
		Content:        template.HTML(mdResult.ContentHTML + "\n" + articleSignSetting),
		Editable:       session.UID == authorModel.ID,
	}

	page := util.GetPage(c)
	commentModels, pagination := service.Comment.GetArticleComments(articleModel.ID, page, blogID)
	var comments []*model.ThemeComment
	for _, commentModel := range commentModels {
		author := &model.ThemeAuthor{}
		if model.SyncCommentAuthorID == commentModel.AuthorID {
			author.URL = commentModel.AuthorURL
			author.Name = commentModel.AuthorName
			author.AvatarURL = commentModel.AuthorAvatarURL
		} else {
			commentAuthor := service.User.GetUser(commentModel.AuthorID)
			commentAuthorBlog := service.User.GetOwnBlog(commentModel.AuthorID)
			author.URL = service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID).Value + util.PathAuthors + "/" + commentAuthor.Name
			author.Name = commentAuthor.Name
			author.AvatarURL = commentAuthor.AvatarURL
		}

		mdResult := util.Markdown(commentModel.Content)
		comment := &model.ThemeComment{
			ID:         commentModel.ID,
			Content:    template.HTML(mdResult.ContentHTML),
			URL:        getBlogURL(c) + articleModel.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(commentModel.ID)),
			Author:     author,
			CreatedAt:  commentModel.CreatedAt.Format("2006-01-02"),
			Removable:  session.UID == authorModel.ID,
			ReplyCount: service.Comment.GetRepliesCount(commentModel.ID, commentModel.BlogID),
		}
		if 0 != commentModel.ParentCommentID {
			parentCommentModel := service.Comment.GetComment(commentModel.ParentCommentID)
			if nil != parentCommentModel {
				parentAuthor := &model.ThemeAuthor{}
				if model.SyncCommentAuthorID == parentCommentModel.AuthorID {
					parentAuthor.URL = parentCommentModel.AuthorURL
					parentAuthor.Name = parentCommentModel.AuthorName
					parentAuthor.AvatarURL = parentCommentModel.AuthorAvatarURL
				} else {
					parentCommentAuthorModel := service.User.GetUser(parentCommentModel.AuthorID)
					commentAuthorBlog := service.User.GetOwnBlog(parentCommentModel.AuthorID)
					parentAuthor.URL = service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID).Value + util.PathAuthors + "/" + parentCommentAuthorModel.Name
					parentAuthor.Name = parentCommentAuthorModel.Name
					parentAuthor.AvatarURL = parentCommentAuthorModel.AvatarURL
				}

				page := service.Comment.GetCommentPage(commentModel.ArticleID, commentModel.ID, commentModel.BlogID)
				parentComment := &model.ThemeComment{
					ID:     parentCommentModel.ID,
					URL:    getBlogURL(c) + articleModel.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(parentCommentModel.ID)),
					Author: parentAuthor,
				}
				comment.Parent = parentComment
			}
		}

		comments = append(comments, comment)
	}

	dataModel["Comments"] = comments
	dataModel["Pagination"] = pagination
	recommendArticleSetting := service.Setting.GetSetting(model.SettingCategoryPreference, model.SettingNamePreferenceRecommendArticleListSize, blogID)
	recommendArticleSize, err := strconv.Atoi(recommendArticleSetting.Value)
	if nil != err {
		recommendArticleSize = 7
	}
	dataModel["RecommendArticles"] = getRecommendArticles(recommendArticleSize)
	fillPreviousArticle(c, articleModel, &dataModel)
	fillNextArticle(c, articleModel, &dataModel)
	dataModel["ToC"] = template.HTML(toc(dataModel["Article"].(*model.ThemeArticle)))
	dataModel["Title"] = articleTitle + " - " + dataModel["Title"].(string)

	c.HTML(http.StatusOK, getTheme(c)+"/article.html", dataModel)

	go service.Article.IncArticleViewCount(articleModel)
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

func getRecommendArticles(size int) []*model.ThemeArticle {
	var ret []*model.ThemeArticle

	if 0 >= size {
		return ret
	}

	indics := gulu.Rand.Ints(0, len(cron.RecommendArticles), size)
	for _, index := range indics {
		article := cron.RecommendArticles[index]

		ret = append(ret, article)
	}

	return ret
}
