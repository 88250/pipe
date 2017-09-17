// Package controller includes controllers.
package controller

import (
	"github.com/gin-gonic/gin"
)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()
	//ret.Use(favicon.New("./favicon.ico"))
	ret.Use(gin.Recovery())

	status := ret.Group("/status")
	{
		status.GET("", statusHandler)
		status.GET("/ping", pingHandler)
	}

	return ret
}
