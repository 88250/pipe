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

// Package i18n includes internationalization related manipulations.
package i18n

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Locale.
type locale struct {
	Name     string
	Langs    map[string]interface{}
	TimeZone string
}

// All locales.
var Locales = map[string]locale{}

// Load loads i18n message configurations.
func Load() {
	f, _ := os.Open("i18n")
	names, _ := f.Readdirnames(-1)
	f.Close()

	if len(Locales) == len(names)-1 {
		return
	}

	for _, name := range names {
		if !strings.HasSuffix(name, ".json") {
			continue
		}

		loc := name[:strings.LastIndex(name, ".")]
		load(loc)
	}
}

func load(localeStr string) {
	bytes, err := ioutil.ReadFile("i18n/" + localeStr + ".json")
	if nil != err {
		log.Error(err)

		os.Exit(-1)
	}

	l := locale{Name: localeStr}

	err = json.Unmarshal(bytes, &l.Langs)
	if nil != err {
		log.Error(err)

		os.Exit(-1)
	}

	Locales[localeStr] = l
}

// Get gets a message with the specified locale and key.
func Get(locale, key string) interface{} {
	return Locales[locale].Langs[key]
}

// GetAll gets all messages with the specified locale.
func GetAll(locale string) map[string]interface{} {
	return Locales[locale].Langs
}

// GetLocalesNames gets names of all locales. Returns ["zh_CN", "en_US"] for example.
func GetLocalesNames() []string {
	ret := []string{}

	for name := range Locales {
		ret = append(ret, name)
	}

	sort.Strings(ret)

	return ret
}
