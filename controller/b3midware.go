package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type DataModel map[string]interface{}

const nilB3id = "H9oxzSym"

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
	if nil != session {
		(*dataModel)["User"] = session
		c.Next()

		return
	} else {
		(*dataModel)["User"] = &util.SessionData{}
	}

	b3id := c.Request.URL.Query().Get("b3id")
	switch b3id {
	case nilB3id:
		c.Next()

		return
	case "":
		redirectURL := strings.TrimSpace(c.Request.Referer())
		if "" == redirectURL {
			redirectURL = util.Conf.Server + c.Request.URL.Path
		}
		c.Redirect(http.StatusSeeOther, "https://hacpai.com/apis/b3-identity?goto="+redirectURL)
		c.Abort()

		return
	default:
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

		session = &util.SessionData{
			UName: username,
			URole: model.UserRoleBlogVisitor,
		}

		user := service.User.GetUserByName(username)
		if nil != user {
			session.BID = user.BlogID
			blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, user.BlogID)
			session.BURL = blogURLSetting.Value
			session.UID = user.ID
			session.URole = user.Role
		}

		if err := session.Save(c); nil != err {
			result.Code = -1
			result.Msg = "saves session failed: " + err.Error()
		}

		(*dataModel)["User"] = session
		c.Next()
	}
}
