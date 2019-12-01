// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
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
