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
	"net/http"
	"strconv"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// UpdateCategoryAction updates a category.
func UpdateCategoryAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}

	category := &model.Category{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(category); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses update category request failed"

		return
	}

	session := util.GetSession(c)
	category.BlogID = session.BID

	if err := service.Category.UpdateCategory(category); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}

// GetCategoryAction gets a category.
func GetCategoryAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}

	data := service.Category.ConsoleGetCategory(id)
	if nil == data {
		result.Code = util.CodeErr

		return
	}

	result.Data = data
}

// GetCategoriesAction gets categories.
func GetCategoriesAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	categoryModels, pagination := service.Category.ConsoleGetCategories(util.GetPage(c), session.BID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)

	var categories []*ConsoleCategory
	for _, categoryModel := range categoryModels {
		categories = append(categories, &ConsoleCategory{
			ID:          categoryModel.ID,
			Title:       categoryModel.Title,
			URL:         blogURLSetting.Value + util.PathCategories + categoryModel.Path,
			Description: categoryModel.Description,
			Number:      categoryModel.Number,
			Tags:        categoryModel.Tags,
		})
	}

	data := map[string]interface{}{}
	data["categories"] = categories
	data["pagination"] = pagination
	result.Data = data
}

// AddCategoryAction adds a category.
func AddCategoryAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)

	category := &model.Category{}
	if err := c.BindJSON(category); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses add category request failed"

		return
	}

	category.BlogID = session.BID
	if err := service.Category.AddCategory(category); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}

// RemoveCategoryAction removes a category.
func RemoveCategoryAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}

	session := util.GetSession(c)
	blogID := session.BID
	if err := service.Category.RemoveCategory(id, blogID); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}
