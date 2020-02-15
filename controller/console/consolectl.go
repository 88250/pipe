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
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// ShowAdminPagesAction shows admin pages.
func ShowAdminPagesAction(c *gin.Context) {
	session := util.GetSession(c)
	if 0 == session.UID {
		c.Redirect(http.StatusSeeOther, model.Conf.Server+"/start")

		return
	}

	t, err := template.ParseFiles(filepath.Join("console/dist/admin" + c.Param("path") + "/index.html"))
	if nil != err {
		logger.Errorf("load console page [" + c.Param("path") + "] failed: " + err.Error())
		c.String(http.StatusNotFound, "load console page failed")

		return
	}

	t.Execute(c.Writer, nil)
}
