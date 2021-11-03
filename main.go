// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/controller"
	"github.com/88250/pipe/cron"
	"github.com/88250/pipe/i18n"
	"github.com/88250/pipe/model"
	"github.com/88250/pipe/service"
	"github.com/88250/pipe/theme"
	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

// Logger
var logger *gulu.Logger

// The only one init function in pipe.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	gulu.Log.SetLevel("debug")
	logger = gulu.Log.NewLogger(os.Stdout)

	model.LoadConf()
	i18n.Load()
	theme.Load()
	replaceServerConf()

	if "dev" == model.Conf.RuntimeMode {
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
		Addr:    "0.0.0.0:" + model.Conf.Port,
		Handler: router,
	}

	go http.ListenAndServe("0.0.0.0:8154", nil)

	handleSignal(server)

	logger.Infof("Pipe (v%s) is running [%s]", util.Version, model.Conf.Server)
	if err := server.ListenAndServe(); nil != err {
		logger.Fatalf("listen and serve failed: " + err.Error())
	}
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logger.Infof("got signal [%s], exiting pipe now", s)
		if err := server.Close(); nil != err {
			logger.Errorf("server close failed: " + err.Error())
		}

		service.DisconnectDB()

		logger.Infof("Pipe exited")
		os.Exit(0)
	}()
}

func replaceServerConf() {
	path := "theme/sw.min.js.tpl"
	if gulu.File.IsExist(path) {
		data, err := ioutil.ReadFile(path)
		if nil != err {
			logger.Fatal("read file [" + path + "] failed: " + err.Error())
		}
		content := string(data)
		content = strings.Replace(content, "http://server.tpl.json", model.Conf.Server, -1)
		content = strings.Replace(content, "http://staticserver.tpl.json", model.Conf.StaticServer, -1)
		content = strings.Replace(content, "${StaticResourceVersion}", model.Conf.StaticResourceVersion, -1)
		writePath := strings.TrimSuffix(path, ".tpl")
		if err = ioutil.WriteFile(writePath, []byte(content), 0644); nil != err {
			logger.Fatal("replace sw.min.js in [" + path + "] failed: " + err.Error())
		}
	}

	if gulu.File.IsExist("console/dist/") {
		err := filepath.Walk("console/dist/", func(path string, f os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".tpl") {
				data, err := ioutil.ReadFile(path)
				if nil != err {
					logger.Fatal("read file [" + path + "] failed: " + err.Error())
				}
				content := string(data)
				content = strings.Replace(content, "http://server.tpl.json", model.Conf.Server, -1)
				content = strings.Replace(content, "http://staticserver.tpl.json", model.Conf.StaticServer, -1)
				writePath := strings.TrimSuffix(path, ".tpl")
				if err = ioutil.WriteFile(writePath, []byte(content), 0644); nil != err {
					logger.Fatal("replace server conf in [" + writePath + "] failed: " + err.Error())
				}
			}

			return err
		})
		if nil != err {
			logger.Fatal("replace server conf in [theme] failed: " + err.Error())
		}
	}
}
