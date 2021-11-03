// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Package model defines variety of utilities.
package model

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/88250/gulu"
	"github.com/88250/pipe/util"
	"github.com/jinzhu/gorm"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// Conf of Pipe.
var Conf *Configuration

// Models represents all models..
var Models = []interface{}{
	&User{}, &Article{}, &Comment{}, &Navigation{}, &Tag{},
	&Category{}, &Archive{}, &Setting{}, &Correlation{},
}

// Table prefix.
const tablePrefix = "b3_pipe_"

// ZeroPushTime represents zero push time.
var ZeroPushTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")

// Configuration (pipe.json).
type Configuration struct {
	Server                string // server scheme, host and port
	StaticServer          string // static resources server scheme, host and port
	StaticResourceVersion string // version of static resources
	LogLevel              string // logging level: trace/debug/info/warn/error/fatal
	ShowSQL               bool   // whether print sql in log
	SessionSecret         string // HTTP session secret
	SessionMaxAge         int    // HTTP session max age (in second)
	RuntimeMode           string // runtime mode (dev/prod)
	SQLite                string // SQLite database file path
	MySQL                 string // MySQL connection URL
	Postgres              string // PostgreSQL connection URL
	Port                  string // listen port
	AxiosBaseURL          string // axios base URL
	MockServer            string // mock server
}

// LoadConf loads the configurations. Command-line arguments will override configuration file.
func LoadConf() {
	version := flag.Bool("version", false, "prints current pipe version")
	confPath := flag.String("conf", "pipe.json", "path of pipe.json")
	confServer := flag.String("server", "", "this will override Conf.Server if specified")
	confStaticServer := flag.String("static_server", "", "this will override Conf.StaticServer if specified")
	confStaticResourceVer := flag.String("static_resource_ver", "", "this will override Conf.StaticResourceVersion if specified")
	confLogLevel := flag.String("log_level", "", "this will override Conf.LogLevel if specified")
	confShowSQL := flag.Bool("show_sql", false, "this will override Conf.ShowSQL if specified")
	confSessionSecret := flag.String("session_secret", "", "this will override Conf.SessionSecret")
	confSessionMaxAge := flag.Int("session_max_age", 0, "this will override Conf.SessionMaxAge")
	confRuntimeMode := flag.String("runtime_mode", "", "this will override Conf.RuntimeMode if specified")
	confSQLite := flag.String("sqlite", "", "this will override Conf.SQLite if specified")
	confMySQL := flag.String("mysql", "", "this will override Conf.MySQL if specified")
	confPostgres := flag.String("postgres", "", "this will override Conf.Postgres if specified")
	confPort := flag.String("port", "", "this will override Conf.Port if specified")
	s2m := flag.Bool("s2m", false, "dumps SQLite data to MySQL SQL script file")

	flag.Parse()

	if *version {
		fmt.Println(util.Version)

		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(*confPath)
	if nil != err {
		logger.Fatal("loads configuration file [" + *confPath + "] failed: " + err.Error())
	}

	Conf = &Configuration{}
	if err = json.Unmarshal(bytes, Conf); nil != err {
		logger.Fatal("parses [pipe.json] failed: ", err)
	}

	gulu.Log.SetLevel(Conf.LogLevel)
	if "" != *confLogLevel {
		Conf.LogLevel = *confLogLevel
		gulu.Log.SetLevel(*confLogLevel)
	}

	if *confShowSQL {
		Conf.ShowSQL = true
	}

	if "" != *confSessionSecret {
		Conf.SessionSecret = *confSessionSecret
	}

	if 0 < *confSessionMaxAge {
		Conf.SessionMaxAge = *confSessionMaxAge
	}

	if "" == Conf.SessionSecret {
		Conf.SessionSecret = gulu.Rand.String(32)
	}

	home, err := gulu.OS.Home()
	if nil != err {
		logger.Fatal("can't find user home directory: " + err.Error())
	}
	logger.Debugf("${home} [%s]", home)

	if "" != *confRuntimeMode {
		Conf.RuntimeMode = *confRuntimeMode
	}

	if "" != *confServer {
		Conf.Server = *confServer
	}

	if "" != *confStaticServer {
		Conf.StaticServer = *confStaticServer
	}
	if "" == Conf.StaticServer {
		Conf.StaticServer = Conf.Server
	}

	time := strconv.FormatInt(time.Now().UnixNano(), 10)
	logger.Debugf("${time} [%s]", time)
	Conf.StaticResourceVersion = strings.Replace(Conf.StaticResourceVersion, "${time}", time, 1)
	if "" != *confStaticResourceVer {
		Conf.StaticResourceVersion = *confStaticResourceVer
	}

	Conf.SQLite = strings.Replace(Conf.SQLite, "${home}", home, 1)
	if "" != *confSQLite {
		Conf.SQLite = *confSQLite
		Conf.MySQL = ""
		Conf.Postgres = ""
	}
	if "" != *confMySQL {
		Conf.MySQL = *confMySQL
		Conf.SQLite = ""
		Conf.Postgres = ""
	}
	if "" != *confPostgres {
		Conf.MySQL = ""
		Conf.SQLite = ""
		Conf.Postgres = *confPostgres
	}

	if "" != *confPort {
		Conf.Port = *confPort
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	if *s2m {
		if "" == Conf.SQLite {
			logger.Fatal("please specify -sqlite")
		}
		if "" == Conf.MySQL {
			logger.Fatal("please specify -mysql")
		}

		sqlite2MySQL(Conf.SQLite, Conf.MySQL)

		os.Exit(0)
	}

	logger.Debugf("configurations [%#v]", Conf)
}
