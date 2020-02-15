// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package service

import (
	"github.com/88250/gulu"
	"log"
	"os"
	"testing"

	"github.com/88250/pipe/model"
)

const (
	testPlatformAdminName = "pipe"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	home, err := gulu.OS.Home()
	if nil != err {
		logger.Fatal(err)
	}

	model.Conf = &model.Configuration{}
	model.Conf.SQLite = home + "/pipe.test.db"

	if gulu.File.IsExist(model.Conf.SQLite) {
		os.Remove(model.Conf.SQLite)
	}

	ConnectDB()

	Init.InitPlatform(&model.User{
		Name:   testPlatformAdminName,
		B3Key:  "beyond",
		Locale: "zh_CN",
	})

	log.Println("setup tests")
}

func teardown() {
	DisconnectDB()

	log.Println("teardown tests")
}
