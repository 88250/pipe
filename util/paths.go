// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"strings"
)

// Path prefixes.
const (
	PathRoot           = "/"
	PathInit           = "/start"
	PathSearch         = "/search"
	PathOpensearch     = "/opensearch.xml"
	PathBlogs          = "/blogs"
	PathConsoleDist    = "/console/dist"
	PathAdmin          = "/admin"
	PathAPI            = "/api"
	PathFavicon        = "/favicon.ico"
	PathTheme          = "/theme"
	PathActivities     = "/activities"
	PathArchives       = "/archives"
	PathArticles       = "/articles"
	PathAuthors        = "/authors"
	PathCategories     = "/categories"
	PathTags           = "/tags"
	PathComments       = "/comments"
	PathAtom           = "/atom"
	PathRSS            = "/rss"
	PathSitemap        = "/sitemap.xml"
	PathChangelogs     = "/changelogs"
	PathRobots         = "/robots.txt"
	PathAPIsSymArticle = "/apis/symphony/article"
	PathPlatInfo       = "/blog/info"
	PathManifest       = "/manifest.json"
)

var reservedPaths = []string{
	PathSearch, PathOpensearch, PathBlogs, PathConsoleDist, PathAdmin, PathAPI, PathFavicon, PathTheme,
	PathActivities, PathArchives, PathAuthors, PathCategories, PathTags, PathComments, PathAtom, PathRSS,
	PathSitemap, PathChangelogs, PathRobots, PathAPIsSymArticle, PathPlatInfo,
}

// IsReservedPath checks the specified path is a reserved path or not.
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
