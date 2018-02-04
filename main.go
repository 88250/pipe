// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/b3log/pipe/controller"
	"github.com/b3log/pipe/cron"
	"github.com/b3log/pipe/i18n"
	"github.com/b3log/pipe/log"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/theme"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

// Logger
var logger *log.Logger

// The only one init function in pipe.
func init() {
	rand.Seed(time.Now().Unix())

	log.SetLevel("warn")
	logger = log.NewLogger(os.Stdout)

	util.LoadConf()
	util.LoadMakrdown()
	i18n.Load()
	theme.Load()
	replaceServerConf()

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
	service.Upgrade.Perform()
	cron.Start()

	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    "0.0.0.0:" + util.Conf.Port,
		Handler: router,
	}

	handleSignal(server)

	logger.Infof("Pipe (v%s) is running [%s]", util.Version, util.Conf.Server)
	server.ListenAndServe()
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logger.Infof("got signal [%s], exiting pipe now", s)
		if err := server.Close(); nil != err {
			logger.Errorf("server close failed: ", err)
		}

		service.DisconnectDB()

		logger.Infof("Pipe exited")
		os.Exit(0)
	}()
}

func replaceServerConf() {
	err := filepath.Walk(filepath.ToSlash(filepath.Join(util.Conf.StaticRoot, "theme")), func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".min.js") {
			data, e := ioutil.ReadFile(path)
			if nil != e {
				logger.Fatal("read file [" + path + "] failed: " + err.Error())
			}
			content := string(data)
			if !strings.Contains(content, "exports={Server:") {
				return err
			}

			json := "{Server:" + strings.Split(content, "{Server:")[1]
			json = strings.Split(json, "}}")[0] + "}"
			newJSON := "{Server:\"" + util.Conf.Server + "\",StaticServer:\"" + util.Conf.StaticServer + "\",StaticResourceVersion:\"" +
				util.Conf.StaticResourceVersion + "\",RuntimeMode:\"" + util.Conf.RuntimeMode + "\",AxiosBaseURL:\"" + util.Conf.AxiosBaseURL +
				"\",MockServer:\"" + util.Conf.MockServer + "\"}"
			content = strings.Replace(content, json, newJSON, -1)
			if e = ioutil.WriteFile(path, []byte(content), 0644); nil != e {
				logger.Fatal("replace server conf in [" + path + "] failed: " + err.Error())
			}
		}

		return err
	})
	if nil != err {
		logger.Fatal("replace server conf in [theme] failed: " + err.Error())
	}

	paths, err := filepath.Glob(filepath.ToSlash(filepath.Join(util.Conf.StaticRoot, "console/dist/*.js")))
	if 0 < len(paths) {
		for _, path := range paths {
			data, e := ioutil.ReadFile(path)
			if nil != e {
				logger.Fatal("read file [" + path + "] failed: " + err.Error())
			}
			content := string(data)
			if strings.Contains(content, "{rel:\"manifest") {
				json := "{rel:\"manifest\",href:\"" + strings.Split(content, "{rel:\"manifest\",href:\"")[1]
				json = strings.Split(json, "}]")[0] + "}"
				newJSON := "{rel:\"manifest\",href:\"" + util.Conf.StaticServer + "/theme/js/manifest.json\"}"
				content = strings.Replace(content, json, newJSON, -1)
			}
			if strings.Contains(content, "env:{Server:") {
				json := "env:{Server:" + strings.Split(content, "env:{Server:")[1]
				json = strings.Split(json, "}}")[0] + "}"
				newJSON := "env:{Server:\"" + util.Conf.Server + "\",StaticServer:\"" + util.Conf.StaticServer + "\",StaticResourceVersion:\"" +
					util.Conf.StaticResourceVersion + "\",RuntimeMode:\"" + util.Conf.RuntimeMode + "\",AxiosBaseURL:\"" + util.Conf.AxiosBaseURL +
					"\",MockServer:\"" + util.Conf.MockServer + "\"}"
				content = strings.Replace(content, json, newJSON, -1)
			}
			if strings.Contains(content, "/console/dist/") {
				part := strings.Split(content, "/console/dist/")[0]
				part = part[strings.LastIndex(part, "\"")+1:]
				content = strings.Replace(content, part, util.Conf.StaticServer, -1)
			}
			if e = ioutil.WriteFile(path, []byte(content), 0644); nil != e {
				logger.Fatal("replace server conf in [" + path + "] failed: " + err.Error())
			}
		}
	}

	if util.File.IsExist("console/dist/") { // dose not exist if npm run dev
		err = filepath.Walk(filepath.ToSlash(filepath.Join(util.Conf.StaticRoot, "console/dist/")), func(path string, f os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".html") {
				data, e := ioutil.ReadFile(path)
				if nil != e {
					logger.Fatal("read file [" + path + "] failed: " + err.Error())
				}
				content := string(data)
				if strings.Contains(content, "rel=\"manifest\" href=\"") {
					rel := "rel=\"manifest\" href=\"" + strings.Split(content, "rel=\"manifest\" href=\"")[1]
					rel = strings.Split(rel, "/>")[0] + "/>"
					newRel := "rel=\"manifest\" href=\"" + util.Conf.StaticServer + "/theme/js/manifest.json\"/>"
					content = strings.Replace(content, rel, newRel, -1)
				}
				if strings.Contains(content, "/console/dist/") {
					part := strings.Split(content, "/console/dist/")[0]
					part = part[strings.LastIndex(part, "\"")+1:]
					content = strings.Replace(content, part, util.Conf.StaticServer, -1)
				}
				v := fmt.Sprintf("%d", time.Now().Unix())
				content = strings.Replace(content, ".js\"", ".js?"+v+"\"", -1)
				content = strings.Replace(content, ".json\"", ".json?"+v+"\"", -1)
				if e = ioutil.WriteFile(path, []byte(content), 0644); nil != e {
					logger.Fatal("replace server conf in [" + path + "] failed: " + err.Error())
				}
			}

			return err
		})
		if nil != err {
			logger.Fatal("replace server conf in [theme] failed: " + err.Error())
		}
	}
}
