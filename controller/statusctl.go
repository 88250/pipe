package controller

import (
	"net/http"

	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func statusHandler(c *gin.Context) {
	result := util.NewResult()
	data := map[string]interface{}{}
	data["articleCount"] = 1

	c.JSON(http.StatusOK, result)
}
