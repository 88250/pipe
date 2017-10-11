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

// Package i18n includes internationalization related manipulations.

package theme

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var themes = []string{}

// Load loads themes.
func Load() {
	f, _ := os.Open("theme/x")
	names, _ := f.Readdirnames(-1)
	f.Close()

	for _, name := range names {
		if !isLetter(rune(name[0])) {
			continue
		}

		themes = append(themes, name)
	}

	log.Debugf("loaded [%d] themes", len(themes))
}

func isLetter(r rune) bool {
	return 'a' <= r && 'z' >= r || 'A' <= r && 'Z' >= r
}
