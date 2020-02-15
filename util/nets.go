// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"github.com/mssola/user_agent"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

// IsDomain checks the specified string is domain name.
func IsDomain(s string) bool {
	return !IsIP(s) && "localhost" != s
}

// IsIP checks the specified string is IP.
func IsIP(s string) bool {
	return nil != net.ParseIP(s)
}

// GetRemoteAddr returns remote address of the context.
func GetRemoteAddr(c *gin.Context) string {
	ret := c.GetHeader("X-forwarded-for")
	ret = strings.TrimSpace(ret)
	if "" == ret {
		ret = c.GetHeader("X-Real-IP")
	}
	ret = strings.TrimSpace(ret)
	if "" == ret {
		return c.Request.RemoteAddr
	}

	return strings.Split(ret, ",")[0]
}

// IsBot checks the specified user-agent is a bot.
func IsBot(uaStr string) bool {
	var ua = user_agent.New(uaStr)

	return ua.Bot() || strings.HasPrefix(uaStr, "Sym")
}
