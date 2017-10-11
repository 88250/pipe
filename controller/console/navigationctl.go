// Solo.go - A small and beautiful blogging platform written in golang.
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
)

type ConsoleNavigation struct {
	ID         uint   `json:"id"`
	Title      string `gorm:"size:128" json:"title"`
	URL        string `json:"url"`
	IconURL    string `json:"iconURL"`
	OpenMethod string `json:"openMethod"`
	Number     int    `json:"number"`
}

func GetNavigationsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)
	navigationModels, pagination := service.Navigation.ConsoleGetNavigations(c.GetInt("p"), sessionData.BID)

	navigations := []*ConsoleNavigation{}
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

func GetNavigationAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1

		return
	}

	data := service.Navigation.ConsoleGetNavigation(uint(id))
	if nil == data {
		result.Code = -1

		return
	}

	result.Data = data
}

func RemoveNavigationAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	if err := service.Navigation.RemoveNavigation(uint(id)); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func UpdateNavigationAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.Atoi(idArg)
	if nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	navigation := &model.Navigation{Model: model.Model{ID: uint(id)}}
	if err := c.BindJSON(navigation); nil != err {
		result.Code = -1
		result.Msg = "parses update navigation request failed"

		return
	}

	if err := service.Navigation.UpdateNavigation(navigation); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}

func AddNavigationAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)

	navigation := &model.Navigation{}
	if err := c.BindJSON(navigation); nil != err {
		result.Code = -1
		result.Msg = "parses add navigation request failed"

		return
	}

	navigation.BlogID = sessionData.BID
	if err := service.Navigation.AddNavigation(navigation); nil != err {
		result.Code = -1
		result.Msg = err.Error()
	}
}
