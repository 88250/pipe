package controller

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func showSearchPageAction(c *gin.Context) {
	t, err := template.ParseFiles("console/dist/search/index.html")
	if nil != err {
		log.Error("load search page failed: " + err.Error())
		c.String(http.StatusNotFound, "load search page failed")

		return
	}

	t.Execute(c.Writer, nil)
}
