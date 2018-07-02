package console

import (
	"net/http"

	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

// UpdatePasswordAction updates a user's password.
func UpdatePasswordAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses update user's password request failed"

		return
	}

	password := arg["password"].(string)

	session := util.GetSession(c)
	user := service.User.GetUserByName(session.UName)
	user.Password = password
	if err := service.User.UpdateUser(user); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}
}

// UpdateAccountAction updates an account.
func UpdateAccountAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses update account request failed"

		return
	}

	b3Key := arg["b3key"].(string)
	avatarURL := arg["avatarURL"].(string)

	session := util.GetSession(c)
	user := service.User.GetUserByName(session.UName)
	user.B3Key = b3Key
	user.AvatarURL = avatarURL
	if err := service.User.UpdateUser(user); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}
	session.UB3Key = b3Key
	session.UAvatar = avatarURL
	session.Save(c)
}

// GetAccountAction gets an account.
func GetAccountAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := util.GetSession(c)
	data := map[string]interface{}{}
	data["name"] = session.UName
	data["avatarURL"] = session.UAvatar
	data["b3Key"] = session.UB3Key

	result.Data = data
}
