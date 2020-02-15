// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

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
