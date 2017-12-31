// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"math"
	"net/http"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func GetTagsAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)

	var tags []*ConsoleTag
	tagModels := service.Tag.GetTags(math.MaxInt64, session.BID)
	for _, tagModel := range tagModels {
		tags = append(tags, &ConsoleTag{Title: tagModel.Title})
	}

	result.Data = tags
}
