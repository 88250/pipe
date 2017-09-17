// Solo.go - A small but beautiful golang blogging system, Solo's golang version.
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
	"os"
	"os/signal"
	"syscall"

	"github.com/b3log/solo.go/controller"
	"github.com/b3log/solo.go/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	f, err := os.Create("solo.log")
	if nil != err {
		panic(err)
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetLevel(log.DebugLevel)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	service.ConnectDB()

	router := controller.MapRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	handleSignal(server)

	log.Info("Solo is running [http://localhost:8080]")

	server.ListenAndServe()

	service.DisconnectDB()
	log.Println("Solo exited")
}

// Configuration (solo.json).
type conf struct {
	IP                    string // server ip, ${ip}
	Port                  string // server port
	Context               string // server context
	Server                string // server host and port ({IP}:{Port})
	StaticServer          string // static resources server scheme, host and port (http://{IP}:{Port})
	StaticResourceVersion string // version of static resources
	LogLevel              string // logging level: debug/info/warn/error
	HTTPSessionMaxAge     int    // HTTP session max age (in seciond)
	RuntimeMode           string // runtime mode (dev/prod)
	WD                    string // current working direcitory, ${pwd}
	Locale                string // default locale
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Infof("Got signal [%s], exiting Solo now", s)
		if err := server.Close(); nil != err {
			log.Fatal("Server Close:", err)
			os.Exit(-1)
		}

		log.Info("Solo exited")
		os.Exit(0)
	}()
}
