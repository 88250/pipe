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

package controller

import (
	"github.com/88250/pipe/service"
	"github.com/gin-gonic/gin"
	"github.com/ikeikeikeike/go-sitemap-generator/stm"
)

func outputSitemapAction(c *gin.Context) {
	sm := stm.NewSitemap(1)
	sm.Create()

	blogs := service.User.GetTopBlogs(10)
	for _, blog := range blogs {
		sm.Add(stm.URL{{"loc", blog.URL}})
	}

	c.Writer.Write(sm.XMLContent())
}
