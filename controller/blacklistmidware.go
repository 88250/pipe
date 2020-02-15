// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
	"github.com/88250/pipe/cron"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func blacklist(c *gin.Context) {
	clientIP := util.GetRemoteAddr(c)

	if _, ok := cron.BlacklistIPs[clientIP]; ok {
		c.Header("Retry-After", "600")
		c.Data(http.StatusTooManyRequests, "text/html; charset=utf-8", []byte("Too Many Requests"))
		c.Abort()
		return
	}

	c.Next()
}
