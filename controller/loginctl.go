package controller

import (
	"net/http"

	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	NameOrEmail    string `json:"nameOrEmail" binding:"required"`
	PasswordHashed string `json:"passwordHashed" binding:"required"`
}

func loginCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	reqData := &loginRequest{}
	if err := c.BindJSON(reqData); nil != err {
		result.Code = -1
		result.Msg = "parses login request failed"

		return
	}

	user := service.User.GetUserByNameOrEmail(reqData.NameOrEmail)
	if nil == user {
		result.Code = -1
		result.Msg = "login failed"

		return
	}

	if user.Password != reqData.PasswordHashed {
		result.Code = -1
		result.Msg = "login failed"

		return
	}

	session := sessions.Default(c)
	session.Set("id", user.ID)
	session.Set("name", user.Name)
	session.Set("role", user.Role)
	session.Save()
}

func logoutCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
}
