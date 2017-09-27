// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
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

package console

import (
	"net/http"
	"strconv"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func AddArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	article := &model.Article{}
	if err := c.BindJSON(article); nil != err {
		result.Code = -1
		result.Msg = "parses add article request failed"

		return
	}

	if err := service.Article.AddArticle(article); nil != err {
		log.Error("add article failed: " + err.Error())
		result.Code = -1
	}
}

func GetArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	data := service.Article.ConsoleGetArticle(uint(id))
	if nil == data {
		result.Code = -1

		return
	}

	result.Data = data
}

func GetArticlesCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	articles, pagination := service.Article.ConsoleGetArticles(c.GetInt("p"))

	data := map[string]interface{}{}
	data["articles"] = articles
	data["pagination"] = pagination
	result.Data = data
}

func RemoveArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	if err := service.Article.RemoveArticle(uint(id)); nil != err {
		log.Error("remove article failed: " + err.Error())
		result.Code = -1
	}
}

func UpdateArticleCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	article := &model.Article{Model: gorm.Model{ID: uint(id)}}
	if err := c.BindJSON(article); nil != err {
		result.Code = -1
		result.Msg = "parses update article request failed"

		return
	}

	if err := service.Article.UpdateArticle(article); nil != err {
		log.Error("update article failed: " + err.Error())
		result.Code = -1
	}
}
