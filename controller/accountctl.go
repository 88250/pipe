package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-contrib/sessions"
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

	ownBlog := service.User.GetOwnBlog(user.ID)
	session := &util.SessionData{
		UID:     user.ID,
		UName:   user.Name,
		UB3Key:  user.B3Key,
		UAvatar: user.AvatarURL,
		URole:   ownBlog.UserRole,
		BID:     ownBlog.ID,
		BURL:    ownBlog.URL,
	}
	if err := session.Save(c); nil != err {
		result.Code = -1
		result.Msg = "saves session failed: " + err.Error()
	}
}

// logoutAction logout a user.
func logoutAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	session.Clear()
	if err := session.Save(); nil != err {
		logger.Errorf("saves session failed: " + err.Error())
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

	if err := service.Init.InitBlog(user); nil != err {
		result.Code = -1
		result.Msg = err.Error()

		return
	}

	ownBlog := service.User.GetOwnBlog(user.ID)
	session := &util.SessionData{
		UID:     user.ID,
		UName:   user.Name,
		UB3Key:  user.B3Key,
		UAvatar: user.AvatarURL,
		URole:   ownBlog.UserRole,
		BID:     ownBlog.ID,
		BURL:    ownBlog.URL,
	}
	if err := session.Save(c); nil != err {
		result.Code = -1
		result.Msg = "saves session failed: " + err.Error()
	}
}

func showLoginPageAction(c *gin.Context) {
	t, err := template.ParseFiles(filepath.ToSlash(filepath.Join(model.Conf.StaticRoot, "console/dist/login/index.html")))
	if nil != err {
		logger.Errorf("load login page failed: " + err.Error())
		c.String(http.StatusNotFound, "load login page failed")

		return
	}

	t.Execute(c.Writer, nil)
}

func showRegisterPageAction(c *gin.Context) {
	t, err := template.ParseFiles(filepath.ToSlash(filepath.Join(model.Conf.StaticRoot, "console/dist/register/index.html")))
	if nil != err {
		logger.Errorf("load register page failed: " + err.Error())
		c.String(http.StatusNotFound, "load register page failed")

		return
	}

	t.Execute(c.Writer, nil)
}
