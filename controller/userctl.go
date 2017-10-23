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
	"net/http"

	"github.com/gin-gonic/gin"
	"strings"
)

func showAuthorsAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	themeAuthorDetail := []*ThemeAuthorDetail{}
	AuthorDetailModels := strings.Split("a, g, c, d", ",")
	for _, authorDetailModel := range AuthorDetailModels {
		authorDetail := &ThemeAuthorDetail{
			Name:      authorDetailModel,
			URL:       "/sss",
			Count:     13,
			AvatarURL: "http://themedesigner.in/demo/admin-press/assets/images/users/2.jpg",
			CreatedAt: "2012-12-12",
		}
		themeAuthorDetail = append(themeAuthorDetail, authorDetail)
	}

	dataModel["Authors"] = themeAuthorDetail
	c.HTML(http.StatusOK, "authors.html", dataModel)
}

func showAuthorArticlesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))
	c.HTML(http.StatusOK, "author-articles.html", dataModel)
}
