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

package service

import (
	"log"
	"os"
	"testing"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	home, err := util.UserHome()
	if nil != err {
		log.Fatal(err)
	}

	util.Conf = &util.Configuration{}
	util.Conf.DataFilePath = home + "/solo.go.test.db"

	ConnectDB()

	log.Println("setup tests")
}

func teardown() {
	DisconnectDB()

	log.Println("teardown tests")
}

func TestInitPlatform(t *testing.T) {
	platformAdmin := &model.User{
		Name:     "sa",
		Email:    "solo.go@b3log.org",
		Password: "saadmin",
		B3Key:    "b3key",
	}

	Init.InitPlatform(platformAdmin)

}
