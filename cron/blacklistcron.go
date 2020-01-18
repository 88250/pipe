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

package cron

import (
	"crypto/tls"
	"github.com/88250/gulu"
	"github.com/88250/pipe/model"
	"github.com/parnurzeal/gorequest"
	"time"
)

// BlacklistIPs saves all banned IPs.
var BlacklistIPs map[string]bool

func refreshBlacklistIPsPeriodically() {
	go refreshBlacklistIPs()

	go func() {
		for range time.Tick(time.Minute * 30) {
			refreshBlacklistIPs()
		}
	}()
}

func refreshBlacklistIPs() {
	defer gulu.Panic.Recover(nil)

	result := map[string]interface{}{}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, _, errs := request.Get("https://hacpai.com/apis/blacklist/ip").
		Set("User-Agent", model.UserAgent).Timeout(3 * time.Second).EndStruct(&result)
	if nil != errs {
		logger.Errorf("refresh blacklist IPs failed: %s", errs)
		return
	}

	BlacklistIPs = map[string]bool{}
	dataIPs := result["data"].([]interface{})
	for _, dataIP := range dataIPs {
		BlacklistIPs[dataIP.(string)] = false
	}
}
