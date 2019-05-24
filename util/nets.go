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