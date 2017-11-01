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
	"strings"

	"github.com/b3log/pipe/i18n"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func addCommentAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	blogAdminVal, _ := c.Get("blogAdmin")
	blogAdmin := blogAdminVal.(*model.User)

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

	dataModel := DataModel{}
	author := &ThemeAuthor{
		Name:      session.UName,
		URL:       "https://hacpai.com/member/" + session.UName,
		AvatarURL: session.UAvatar,
	}
	themeComment := ThemeComment{
		ID:        comment.ID,
		Content:   template.HTML(util.Markdown(comment.Content)),
		Author:    author,
		CreatedAt: comment.CreatedAt.Format("2006-01-02"),
		Removable: false,
	}
	dataModel["Item"] = themeComment
	dataModel["ArticleID"] = comment.ArticleID

	localeSetting := service.Setting.GetSetting(model.SettingCategoryI18n, model.SettingNameI18nLocale, blogAdmin.BlogID)
	i18ns := i18n.GetMessages(localeSetting.Value)
	i18nMap := map[string]interface{}{}
	for key, value := range i18ns {
		i18nMap[strings.Title(key)] = value
		i18nMap[key] = value
	}
	dataModel["I18n"] = i18nMap

	t := template.Must(template.New("").ParseFiles("theme/x/" + getTheme(c) + "/define-comment.html"))

	htmlBuilder := bytes.Buffer{}
	if err := t.ExecuteTemplate(&htmlBuilder, getTheme(c)+"/comment", dataModel); nil != err {
		log.Errorf("execute comment template failed: " + err.Error())

		return
	}

	result.Data = htmlBuilder.String()
}
