// Solo.go - A small and beautiful blogging platform written in golang.
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

// Package util defines variety of utilities.
package util

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Solo configuration.
var Conf *Configuration

// Configuration (solo.json).
type Configuration struct {
	Host                  string // server host
	Port                  string // server port
	Context               string // server context
	Server                string // server host and port ({IP}:{Port})
	StaticServer          string // static resources server scheme, host and port (http://{IP}:{Port})
	StaticResourceVersion string // version of static resources
	LogFilePath           string // log file path
	LogLevel              string // logging level: debug/info/warn/error/fatal
	SessionSecret         string // HTTP session secret
	SessionMaxAge         int    // HTTP session max age (in seciond)
	RuntimeMode           string // runtime mode (dev/prod)
	WD                    string // current working direcitory, ${pwd}
	DataFilePath          string // database file path
}

// InitConf initializes the conf. Args will override configuration file.
func InitConf() {
	confPath := flag.String("conf", "solo.json", "path of solo.json")
	confHost := flag.String("host", "", "this will override Solo.Host if specified")
	confPort := flag.String("port", "", "this will override Solo.Port if specified")
	confContext := flag.String("context", "", "this will override Solo.Context if specified")
	confServer := flag.String("server", "", "this will override Solo.Server if specified")
	confStaticServer := flag.String("static_server", "", "this will override Solo.StaticServer if specified")
	confStaticResourceVer := flag.String("static_resource_ver", "", "this will override Solo.StaticResourceVersion if specified")
	confLogFilePath := flag.String("log_file_path", "", "this will override Solo.LogFilePath if specified")
	confLogLevel := flag.String("log_level", "", "this will override Solo.LogLevel if specified")
	confDataFilePath := flag.String("data_file_path", "", "this will override Solo.DataFilePath if specified")

	flag.Parse()

	bytes, err := ioutil.ReadFile(*confPath)
	if nil != err {
		log.Fatal("loads configuration file [" + *confPath + "] failed: " + err.Error())
	}

	Conf = &Configuration{}
	if err = json.Unmarshal(bytes, Conf); nil != err {
		log.Fatal("parses [solo.json] failed: ", err)
	}

	home, err := UserHome()
	if nil != err {
		log.Fatal("can't find user home directory: " + err.Error())
	}
	Conf.LogFilePath = strings.Replace(Conf.LogFilePath, "${home}", home, 1)
	if "" != *confLogFilePath {
		Conf.LogFilePath = *confLogFilePath
	}
	f, err := os.OpenFile(Conf.LogFilePath, os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		log.Fatal("creates log file [" + Conf.LogFilePath + "] failed: " + err.Error())
	}
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	log.SetLevel(getLogLevel(Conf.LogLevel))
	if "" != *confLogLevel {
		Conf.LogLevel = *confLogLevel
		log.SetLevel(getLogLevel(*confLogLevel))
	}
	log.Debugf("${home} [%s]", home)

	if "" != *confHost {
		Conf.Host = *confHost
	}

	if "" != *confPort {
		Conf.Port = *confPort
	}

	if "" != *confContext {
		Conf.Context = *confContext
	}

	Conf.Server = strings.Replace(Conf.Server, "{Host}", Conf.Host, 1)
	Conf.Server = strings.Replace(Conf.Server, "{Port}", Conf.Port, 1)
	if "" != *confServer {
		Conf.Server = *confServer
	}

	Conf.StaticServer = strings.Replace(Conf.StaticServer, "{Host}", Conf.Host, 1)
	Conf.StaticServer = strings.Replace(Conf.StaticServer, "{Port}", Conf.Port, 1)
	if "" != *confStaticServer {
		Conf.StaticServer = *confStaticServer
	}

	time := strconv.FormatInt(time.Now().UnixNano(), 10)
	log.Debugf("${time} [%s]", time)
	Conf.StaticResourceVersion = strings.Replace(Conf.StaticResourceVersion, "${time}", time, 1)
	if "" != *confStaticResourceVer {
		Conf.StaticResourceVersion = *confStaticResourceVer
	}

	Conf.DataFilePath = strings.Replace(Conf.DataFilePath, "${home}", home, 1)
	if "" != *confDataFilePath {
		Conf.DataFilePath = *confDataFilePath
	}

	log.Debugf("configurations [%+v]", Conf)
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
