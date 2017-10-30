// Pipe - A small and beautiful blogging platform written in golang.
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
	PathRoot       = "/"
	PathInit       = "/init"
	PathSearch     = "/search"
	PathBlogs      = "/blogs"
	PathAssets     = "/assets"
	PathAdmin      = "/admin"
	PathAPI        = "/api"
	PathFavicon    = "/favicon.ico"
	PathTheme      = "/theme"
	PathActivities = "/activities"
	PathArchives   = "/archives"
	PathArticles   = "/articles"
	PathAuthors    = "/authors"
	PathCategories = "/categories"
	PathTags       = "/tags"
	PathComments   = "/comments"
)

var reservedPaths = []string{
	PathInit, PathSearch, PathBlogs, PathAssets, PathAdmin, PathAPI,
	PathFavicon, PathTheme, PathActivities, PathArchives, PathArticles, PathAuthors,
	PathCategories, PathTags, PathComments,
}

func IsReservedPath(path string) bool {
	path = strings.TrimSpace(path)
	if PathRoot == path {
		return true
	}

	for _, reservedPath := range reservedPaths {
		if strings.HasPrefix(path, reservedPath) {
			return true
		}
	}

	return false
}
