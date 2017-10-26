package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func fillUser(c *gin.Context) {
	inited := service.Init.Inited()
	if !inited && util.PathInit != c.Request.URL.Path {
		c.Redirect(http.StatusSeeOther, util.Conf.Server+util.PathInit)

		c.Abort()

		return
	}

	dataModel := &DataModel{}
	c.Set("dataModel", dataModel)

	session := util.GetSession(c)
	if nil == session {
		session = &util.SessionData{}
	}
	(*dataModel)["user"] = session

	c.Next()
}

func b3IdCheck(c *gin.Context) {
	b3id := c.Request.URL.Query().Get("b3id")
	if "" == b3id {
		c.Next()

		return
	}

	httpClient := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	res, err := httpClient.Get("https://hacpai.com/apis/check-b3-identity?b3id=" + b3id)
	if nil != err {
		log.Error("check b3 identity failed: " + err.Error())
		c.Next()

		return
	}
	defer res.Body.Close()

	result := util.NewResult()
	if err := json.NewDecoder(res.Body).Decode(result); nil != err {
		log.Error("parse b3 identity check result failed: " + err.Error())
		c.Next()

		return
	}

	if 0 != result.Code {
		c.Next()

		return
	}

	data := result.Data.(map[string]interface{})
	username := data["userName"].(string)

	sessionData := &util.SessionData{
		UName: username,
		URole: model.UserRoleBlogVisitor,
	}

	user := service.User.GetUserByName(username)
	if nil != user {
		sessionData.BID = user.BlogID
		blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, user.BlogID)
		sessionData.BURL = blogURLSetting.Value
		sessionData.UID = user.ID
		sessionData.URole = user.Role
	}

	if err := sessionData.Save(c); nil != err {
		result.Code = -1
		result.Msg = "saves session failed: " + err.Error()
	}

	c.Next()
}
