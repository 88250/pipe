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
	"github.com/b3log/pipe/util"
	"strings"
)

func showCategoriesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))

	themeCategoryDetail := []*ThemeCategoryDetail{}
	CategoriesModels := strings.Split("a, g, c, d", ",")
	for _, categoriesModel := range CategoriesModels {

		themeTags := []*ThemeTag{}
		tagStrs := strings.Split("a, g, c, d", ",")
		for _, tagStr := range tagStrs {
			themeTag := &ThemeTag{
				Title: tagStr,
				URL:   getSystemPath(c) + util.PathTags + "/" + tagStr,
			}
			themeTags = append(themeTags, themeTag)
		}

		categoriesDetail := &ThemeCategoryDetail{
			Title:       categoriesModel,
			URL:         "/sss",
			Description: "http://themedesigner.in/demo/admin-press/assets/images/users/2.jpg",
			Tags:        themeTags,
			Count:       23,
		}
		themeCategoryDetail = append(themeCategoryDetail, categoriesDetail)
	}

	dataModel["Categories"] = themeCategoryDetail
	c.HTML(http.StatusOK, "categories.html", dataModel)
}

func showCategoryArticlesArticlesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))
	c.HTML(http.StatusOK, "category-articles.html", dataModel)
}
