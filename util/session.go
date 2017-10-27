// Pipe - A small and beautiful blogging platform written in golang.
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

package util

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	UID   uint   // user ID
	UName string // username
	URole int    // user role
	BID   uint   // blog ID
	BURL  string // blog url
}

func (sd *SessionData) Save(c *gin.Context) error {
	session := sessions.Default(c)
	sessionDataBytes, err := json.Marshal(sd)
	if nil != err {
		return err
	}
	session.Set("data", string(sessionDataBytes))

	return session.Save()
}

func GetSession(c *gin.Context) *SessionData {
	session := sessions.Default(c)
	sessionDataStr := session.Get("data")
	if nil == sessionDataStr {
		return nil
	}

	ret := &SessionData{}
	err := json.Unmarshal([]byte(sessionDataStr.(string)), ret)
	if nil != err {
		return nil
	}

	c.Set("session", ret)

	return ret
}
