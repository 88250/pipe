// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/b3log/solo.go/controller"
	"github.com/b3log/solo.go/i18n"
	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/theme"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// The only one init function in Solo.go.
func init() {
	util.LoadConf()
	i18n.Load()
	theme.Load()

	if "dev" == util.Conf.RuntimeMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
}

// Entry point.
func main() {
	service.ConnectDB()

	serverURL, err := url.ParseRequestURI(util.Conf.Server)
	if nil != err {
		log.Fatal("Invalid [Server] configuration item")
	}

	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    serverURL.Host,
		Handler: router,
	}

	handleSignal(server)

	log.Infof("Solo.go (v%s) is running [%s]", util.Version, util.Conf.Server)

	server.ListenAndServe()
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Infof("got signal [%s], exiting Solo.go now", s)
		if err := server.Close(); nil != err {
			log.Error("server close failed: ", err)
		}

		service.DisconnectDB()

		log.Info("Solo exited")
		os.Exit(0)
	}()
}
