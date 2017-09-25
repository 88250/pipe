// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
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

package controller

import (
	"net/http"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type initRequest struct {
	name     string `json:"name" binding:"required"`
	email    string `json:"email" binding:"required"`
	password string `json:"password" binding:"required"`
}

func initHandler(c *gin.Context) {
	result := util.NewResult()

	reqData := &initRequest{}
	if err := c.BindJSON(&reqData); nil != err {
		log.Error("parses init request failed: " + err.Error())

		result.Code = -1
		result.Msg = "parses init request fialed"

		return
	}

	platformAdmin := &model.User{
		Name:     reqData.name,
		Email:    reqData.email,
		Password: reqData.password,
	}

	if err := service.Init.InitPlatform(platformAdmin); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	c.JSON(http.StatusOK, result)
}
