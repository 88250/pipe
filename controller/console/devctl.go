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

package console

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// GenArticlesAction generates articles for testing.
func GenArticlesAction(c *gin.Context) {
	session := util.GetSession(c)

	for i := 0; i < 100; i++ {
		article := &model.Article{
			AuthorID: session.UID,
			Title:    "title " + strconv.Itoa(i) + "_" + strconv.Itoa(rand.Int()),
			Tags:     "开发生成",
			Content:  "开发生成",
			BlogID:   session.BID,
		}
		if err := service.Article.AddArticle(article); nil != err {
			logger.Errorf("generate article failed: " + err.Error())
		}
	}

	c.Redirect(http.StatusTemporaryRedirect, model.Conf.Server)
}
