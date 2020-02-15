// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Package i18n includes internationalization related manipulations.
package i18n

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/88250/gulu"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

type locale struct {
	Name     string
	Langs    map[string]interface{}
	TimeZone string
}

var locales = map[string]locale{}

// Load loads i18n message configurations.
func Load() {
	f, _ := os.Open("i18n")
	names, _ := f.Readdirnames(-1)
	f.Close()

	for _, name := range names {
		if !gulu.Rune.IsLetter(rune(name[0])) || !strings.HasSuffix(name, ".json") {
			continue
		}

		loc := name[:strings.LastIndex(name, ".")]
		load(loc)
	}

	logger.Tracef("loaded [%d] language configuration files", len(locales))
}

func load(localeStr string) {
	bytes, err := ioutil.ReadFile("i18n/" + localeStr + ".json")
	if nil != err {
		logger.Fatal("reads i18n configurations fialed: " + err.Error())
	}

	l := locale{Name: localeStr}

	err = json.Unmarshal(bytes, &l.Langs)
	if nil != err {
		logger.Fatal("parses i18n configurations failed: " + err.Error())
	}

	locales[localeStr] = l
}

// GetMessagef gets a message with the specified locale, key and arguments
func GetMessagef(locale, key string, a ...interface{}) string {
	msg := GetMessage(locale, key)

	return fmt.Sprintf(msg, a...)
}

// GetMessage gets a message with the specified locale and key.
func GetMessage(locale, key string) string {
	return locales[locale].Langs[key].(string)
}

// GetMessages gets all messages with the specified locale.
func GetMessages(locale string) map[string]interface{} {
	return locales[locale].Langs
}

// GetLocalesNames gets names of all locales. Returns ["zh_CN", "en_US"] for example.
func GetLocalesNames() []string {
	var ret []string

	for name := range locales {
		ret = append(ret, name)
	}

	sort.Strings(ret)

	return ret
}
