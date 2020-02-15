// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Package controller is the "controller" layer.
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showActivitiesAction(c *gin.Context) {
	dm, _ := c.Get("dataModel")
	dataModel := *(dm.(*DataModel))
	c.HTML(http.StatusOK, getTheme(c)+"/activities.html", dataModel)
}
