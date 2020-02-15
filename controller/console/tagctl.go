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
	"math"
	"net/http"
	"strconv"

	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// GetTagsAction gets tags.
func GetTagsAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)

	var tags []*ConsoleTag
	tagModels := service.Tag.GetTags(math.MaxInt64, session.BID)
	for _, tagModel := range tagModels {
		tags = append(tags, &ConsoleTag{
			Title: tagModel.Title,
			URL:   blogURLSetting.Value + util.PathTags + "/" + tagModel.Title,
		})
	}

	result.Data = tags
}

// GetTagsAction gets tags with pagination.
func GetTagsPageAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	tagModels, pagination := service.Tag.ConsoleGetTags(c.Query("key"), util.GetPage(c), session.BID)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, session.BID)

	var tags []*ConsoleTag
	for _, tagModel := range tagModels {
		item := &ConsoleTag{
			ID:    tagModel.ID,
			Title: tagModel.Title,
			URL:   blogURLSetting.Value + util.PathTags + "/" + tagModel.Title,
		}
		tags = append(tags, item)
	}
	data := map[string]interface{}{}
	data["tags"] = tags
	data["pagination"] = pagination
	result.Data = data
}

// RemoveTagsAction remove tags that have no articles.
func RemoveTagsAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}
	session := util.GetSession(c)
	blogID := session.BID

	if err := service.Tag.RemoveTag(id, blogID); nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()
	}

}
