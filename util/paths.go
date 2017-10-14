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

package util

import (
	"strings"
)

// Path prefixes.
const (
	pathAssets     = "/assets"
	pathAdmin      = "/admin"
	pathAPI        = "/api"
	pathFavicon    = "/favicon.ico"
	pathTheme      = "/theme"
	pathActivities = "/activities"
	pathArchives   = "/archives"
	pathArticles   = "/articles"
	pathAuthors    = "/authors"
	pathCategories = "/categories"
	pathTags       = "/tags"
)

var reservedPaths = []string{
	pathAssets, pathAdmin, pathAPI, pathFavicon, pathTheme, pathActivities,
	pathArchives, pathArticles, pathAuthors, pathCategories, pathTags,
}

func IsReservedPath(path string) bool {
	for _, reservedPath := range reservedPaths {
		if strings.HasPrefix(path, reservedPath) {
			return true
		}
	}

	return false
}
