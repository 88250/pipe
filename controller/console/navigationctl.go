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

// GetNavigationsAction gets navigations.
func GetNavigationsAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	navigationModels, pagination := service.Navigation.ConsoleGetNavigations(util.GetPage(c), session.BID)

	var navigations []*ConsoleNavigation
	for _, navigationModel := range navigationModels {
		comment := &ConsoleNavigation{
			ID:         navigationModel.ID,
			Title:      navigationModel.Title,
			URL:        navigationModel.URL,
			IconURL:    navigationModel.IconURL,
			OpenMethod: navigationModel.OpenMethod,
			Number:     navigationModel.Number,
		}

		navigations = append(navigations, comment)
	}

	data := map[string]interface{}{}
	data["navigations"] = navigations
	data["pagination"] = pagination
	result.Data = data
}

// GetNavigationAction gets a navigation.
func GetNavigationAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr

		return
	}

	data := service.Navigation.ConsoleGetNavigation(uint64(id))
	if nil == data {
		result.Code = util.CodeErr

		return
	}

	result.Data = data
}

// RemoveNavigationAction remove a navigation.
func RemoveNavigationAction(c *gin.Context) {
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

	if err := service.Navigation.RemoveNavigation(uint64(id), blogID); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}

// UpdateNavigationAction updates a navigation.
func UpdateNavigationAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}

	navigation := &model.Navigation{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(navigation); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses update navigation request failed"

		return
	}

	session := util.GetSession(c)
	navigation.BlogID = session.BID

	if err := service.Navigation.UpdateNavigation(navigation); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}

// AddNavigationAction adds a navigation.
func AddNavigationAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)

	navigation := &model.Navigation{}
	if err := c.BindJSON(navigation); nil != err {
		result.Code = util.CodeErr
		result.Msg = "parses add navigation request failed"

		return
	}

	navigation.BlogID = session.BID
	if err := service.Navigation.AddNavigation(navigation); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}
}
