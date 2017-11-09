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

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func getRepliesAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogAdmin := getBlogAdmin(c)
	parentCmtIDArg := strings.SplitAfter(c.Request.URL.Path, util.PathComments+"/")[1]
	parentCmtIDArg = strings.Split(parentCmtIDArg, "/replies")[0]
	parentCmtID, _ := strconv.Atoi(parentCmtIDArg)

	replyComments := service.Comment.GetReplies(uint(parentCmtID), blogAdmin.BlogID)
	replies := []*ThemeReply{}
	for _, replyComment := range replyComments {
		commentAuthor := service.User.GetUser(replyComment.AuthorID)
		if nil == commentAuthor {
			logger.Errorf("not found comment author [userID=%d]", replyComment.AuthorID)

			continue
		}
		commentAuthorURL := util.HacPaiURL + "/member/" + commentAuthor.Name
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, commentAuthor.BlogID)
		if nil != blogURLSetting {
			commentAuthorURL = blogURLSetting.Value + util.PathAuthors + "/" + commentAuthor.Name
		}

		author := &ThemeAuthor{
			Name:      commentAuthor.Name,
			URL:       commentAuthorURL,
			AvatarURL: commentAuthor.AvatarURLWithSize(64),
		}

		reply := &ThemeReply{
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
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogAdmin := getBlogAdmin(c)
	session := util.GetSession(c)
	if nil == session {
		result.Code = -1
		result.Msg = "Please login before comment"

		return
	}

	comment := &model.Comment{
		AuthorID: session.UID,
		BlogID:   blogAdmin.BlogID,
	}
	if err := c.BindJSON(comment); nil != err {
		result.Code = -1
		result.Msg = "parses add comment request failed"

		return
	}

	if err := service.Comment.AddComment(comment); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}

	dataModel := getDataModel(c)

	commentAuthorURL := util.HacPaiURL + "/member/" + session.UName
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)
	if nil != blogURLSetting {
		commentAuthorURL = blogURLSetting.Value + util.PathAuthors + "/" + session.UName
	}
	author := &ThemeAuthor{
		Name:      session.UName,
		URL:       commentAuthorURL,
		AvatarURL: session.UAvatar,
	}
	themeComment := ThemeComment{
		ID:        comment.ID,
		Content:   template.HTML(util.Markdown(comment.Content).ContentHTML),
		Author:    author,
		CreatedAt: comment.CreatedAt.Format("2006-01-02"),
		Removable: false,
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
