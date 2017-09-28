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

	"github.com/b3log/solo.go/service"
	"github.com/b3log/wide/util"
	"github.com/gin-gonic/gin"
)

func GetTagsCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	tags := []*TagPermalink{}
	tagModels := service.Tag.ConsoleGetTags()
	for _, tagModel := range tagModels {
		tag := &TagPermalink{
			Title:     tagModel.Title,
			Permalink: "context/" + tagModel.Title,
		}

		tags = append(tags, tag)
	}

	result.Data = tags
}
