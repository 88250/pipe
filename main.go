// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
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
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/b3log/solo.go/controller"
	"github.com/b3log/solo.go/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Solo configuration.
var Conf *conf

// The only one init function in Solo.
func init() {
	confPath := flag.String("conf", "solo.json", "path of solo.json")
	confLogPath := flag.String("log_path", "solo.log", "path of log file")
	confHost := flag.String("host", "", "this will override Solo.Host if specified")
	confPort := flag.String("port", "", "this will override Solo.Port if specified")
	confContext := flag.String("context", "", "this will override Solo.Context if specified")
	confServer := flag.String("server", "", "this will override Solo.Server if specified")
	confStaticServer := flag.String("static_server", "", "this will override Solo.StaticServer if specified")
	confStaticResourceVer := flag.String("static_resource_ver", "", "this will override Solo.StaticResourceVersion if specified")
	confLogLevel := flag.String("log_level", "", "this will override Solo.LogLevel if specified")

	flag.Parse()

	args := map[string]interface{}{}
	args["confPath"] = *confPath
	args["confLogPath"] = *confLogPath
	args["confHost"] = *confHost
	args["confPort"] = *confPort
	args["confContext"] = *confContext
	args["confServer"] = *confServer
	args["confStaticServer"] = *confStaticServer
	args["confStaticResourceVer"] = *confStaticResourceVer
	args["confLogLevel"] = *confLogLevel

	f, err := os.Create(*confLogPath)
	if nil != err {
		log.Fatal("Creates log file [" + *confLogPath + "] failed: " + err.Error())
	}
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	initConf(&args)
}

// Entry point.
func main() {
	service.ConnectDB()

	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    Conf.Server,
		Handler: router,
	}

	handleSignal(server)

	log.Infof("Solo is running [%s]", Conf.Server)

	server.ListenAndServe()
}

// Configuration (solo.json).
type conf struct {
	Host                  string // server host
	Port                  string // server port
	Context               string // server context
	Server                string // server host and port ({IP}:{Port})
	StaticServer          string // static resources server scheme, host and port (http://{IP}:{Port})
	StaticResourceVersion string // version of static resources
	LogLevel              string // logging level: debug/info/warn/error/fatal
	HTTPSessionMaxAge     int    // HTTP session max age (in seciond)
	RuntimeMode           string // runtime mode (dev/prod)
	WD                    string // current working direcitory, ${pwd}
	Locale                string // default locale
}

// initConf initializes the conf. Args will over
func initConf(args *map[string]interface{}) {
	confs := *args
	confPath := confs["confPath"].(string)
	bytes, err := ioutil.ReadFile(confPath)
	if nil != err {
		log.Fatal("loads configuration file [" + confPath + "] failed: " + err.Error())
	}

	Conf = &conf{}
	if err = json.Unmarshal(bytes, Conf); nil != err {
		log.Fatal("parses [solo.json] failed: ", err)
	}

	log.SetLevel(getLogLevel(Conf.LogLevel))
	if confLogLevel := confs["confLogLevel"].(string); "" != confLogLevel {
		Conf.LogLevel = confLogLevel
		log.SetLevel(getLogLevel(confLogLevel))
	}

	if confHost := confs["confHost"].(string); "" != confHost {
		Conf.Host = confHost
	}

	if confPort := confs["confPort"].(string); "" != confPort {
		Conf.Port = confPort
	}

	if confContext := confs["confContext"].(string); "" != confContext {
		Conf.Context = confContext
	}

	Conf.Server = strings.Replace(Conf.Server, "{Host}", Conf.Host, 1)
	Conf.Server = strings.Replace(Conf.Server, "{Port}", Conf.Port, 1)
	if confServer := confs["confServer"].(string); "" != confServer {
		Conf.Server = confServer
	}

	Conf.StaticServer = strings.Replace(Conf.StaticServer, "{Host}", Conf.Host, 1)
	Conf.StaticServer = strings.Replace(Conf.StaticServer, "{Port}", Conf.Port, 1)
	if confStaticServer := confs["confStaticServer"].(string); "" != confStaticServer {
		Conf.StaticServer = confStaticServer
	}

	time := strconv.FormatInt(time.Now().UnixNano(), 10)
	log.Debugf("${time} [%s]", time)
	Conf.StaticResourceVersion = strings.Replace(Conf.StaticResourceVersion, "${time}", time, 1)
	if confStaticResourceVer := confs["confStaticResourceVer"].(string); "" != confStaticResourceVer {
		Conf.StaticResourceVersion = confStaticResourceVer
	}

	log.Debugf("Conf [%+v]", Conf)
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		log.Infof("Got signal [%s], exiting Solo now", s)
		if err := server.Close(); nil != err {
			log.Error("server close failed: ", err)
		}

		service.DisconnectDB()

		log.Info("Solo exited")
		os.Exit(0)
	}()
}

// getLogLevel gets logging level value (logrus.level) corresponding to the specified level.
func getLogLevel(level string) log.Level {
	level = strings.ToLower(level)

	switch level {
	case "trace", "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}
