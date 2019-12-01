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

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

func getRepliesAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogID := getBlogID(c)
	parentCmtIDArg := strings.SplitAfter(c.Request.URL.Path, util.PathComments+"/")[1]
	parentCmtIDArg = strings.Split(parentCmtIDArg, "/replies")[0]
	parentCmtID, _ := strconv.ParseUint(parentCmtIDArg, 10, 64)

	replyComments := service.Comment.GetReplies(parentCmtID, blogID)
	var replies []*model.ThemeReply
	for _, replyComment := range replyComments {
		commentAuthor := service.User.GetUser(replyComment.AuthorID)
		if nil == commentAuthor {
			logger.Errorf("not found comment author [userID=%d]", replyComment.AuthorID)

			continue
		}
		commentAuthorBlog := service.User.GetOwnBlog(commentAuthor.ID)
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthorBlog.ID)
		commentAuthorURL := blogURLSetting.Value + util.PathAuthors + "/" + commentAuthor.Name
		author := &model.ThemeAuthor{
			Name:      commentAuthor.Name,
			URL:       commentAuthorURL,
			AvatarURL: commentAuthor.AvatarURLWithSize(64),
		}

		reply := &model.ThemeReply{
			ID:        replyComment.ID,
			Content:   template.HTML(util.Markdown(replyComment.Content).ContentHTML),
			Author:    author,
			CreatedAt: replyComment.CreatedAt.Format("2006-01-02"),
		}
		replies = append(replies, reply)
	}

	result.Data = replies
}

func addCommentAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogID := getBlogID(c)
	session := util.GetSession(c)
	if 0 == session.UID {
		result.Code = util.CodeErr
		result.Msg = "please login before comment"

		return
	}

	comment := &model.Comment{
		AuthorID: session.UID,
		BlogID:   blogID,
	}
	if err := c.BindJSON(comment); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses add comment request failed"

		return
	}

	comment.IP = util.GetRemoteAddr(c)

	if err := service.Comment.AddComment(comment); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}

	dataModel := getDataModel(c)

	commentAuthorURL := util.HacPaiURL + "/member/" + session.UName
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)
	if nil != blogURLSetting {
		commentAuthorURL = blogURLSetting.Value + util.PathAuthors + "/" + session.UName
	}
	author := &model.ThemeAuthor{
		Name:      session.UName,
		URL:       commentAuthorURL,
		AvatarURL: session.UAvatar,
	}
	page := service.Comment.GetCommentPage(comment.ArticleID, comment.ID, comment.BlogID)
	article := service.Article.ConsoleGetArticle(comment.ArticleID)
	themeComment := &model.ThemeComment{
		ID:        comment.ID,
		Content:   template.HTML(util.Markdown(comment.Content).ContentHTML),
		URL:       getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(comment.ID)),
		Author:    author,
		CreatedAt: comment.CreatedAt.Format("2006-01-02"),
		Removable: false,
	}
	if 0 != comment.ParentCommentID {
		parentCommentModel := service.Comment.GetComment(comment.ParentCommentID)
		if nil != parentCommentModel {
			parentCommentAuthorName := parentCommentModel.AuthorName
			if "" == parentCommentAuthorName {
				parentCommentAuthorModel := service.User.GetUser(parentCommentModel.AuthorID)
				parentCommentAuthorName = parentCommentAuthorModel.Name
			}
			parentComment := &model.ThemeComment{
				ID:  parentCommentModel.ID,
				URL: getBlogURL(c) + article.Path + "?p=" + strconv.Itoa(page) + "#pipeComment" + strconv.Itoa(int(parentCommentModel.ID)),
				Author: &model.ThemeAuthor{
					Name: parentCommentAuthorName,
				},
			}
			themeComment.Parent = parentComment
		}
	}
	dataModel["Item"] = themeComment
	dataModel["ArticleID"] = comment.ArticleID

	t := template.Must(template.New("").ParseFiles("theme/comment/comment.html"))

	htmlBuilder := bytes.Buffer{}
	if err := t.ExecuteTemplate(&htmlBuilder, "comment/comment", dataModel); nil != err {
		logger.Errorf("execute comment template failed: " + err.Error())

		return
	}

	result.Data = htmlBuilder.String()
}
