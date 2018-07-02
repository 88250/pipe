package controller

import (
	"net/http"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

// loginAction login a user.
func loginAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses login request failed"

		return
	}

	name := arg["name"].(string)
	password := arg["password"].(string)

	user := service.User.GetUserByName(name)
	if nil == user {
		result.Code = -1
		result.Msg = "user not found"

		return
	}

	crypt := sha512_crypt.New()
	inputHash, _ := crypt.Generate([]byte(password), []byte(user.Password))
	if inputHash != user.Password {
		result.Code = -1
		result.Msg = "wrong password"

		return
	}
}

// registerAction registers a user.
func registerAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Code = -1
		result.Msg = "parses register request failed"

		return
	}

	name := arg["name"].(string)
	password := arg["password"].(string)

	existUser := service.User.GetUserByName(name)
	if nil != existUser {
		result.Code = -1
		result.Msg = "duplicated user name"

		return
	}

	user := &model.User{
		Name:      name,
		Password:  password,
		AvatarURL: "https://img.hacpai.com/pipe/default-avatar.png",
	}

	if err := service.User.AddUser(user); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}
}
