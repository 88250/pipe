// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package console

import (
	"html/template"

	"github.com/88250/pipe/util"
)

// ConsoleArticle represents console article.
type ConsoleArticle struct {
	ID           uint64         `json:"id"`
	Author       *ConsoleAuthor `json:"author"`
	CreatedAt    string         `json:"createdAt"`
	Title        string         `json:"title"`
	Tags         []*ConsoleTag  `json:"tags"`
	URL          string         `json:"url"`
	Topped       bool           `json:"topped"`
	ViewCount    int            `json:"viewCount"`
	CommentCount int            `json:"commentCount"`
}

// ConsoleTag represents console tag.
type ConsoleTag struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url,omitempty"`
}

// ConsoleAuthor represents console author.
type ConsoleAuthor struct {
	URL       string `json:"url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarURL"`
}

// AvatarURLWithSize returns avatar URL with the specified size.
func (u *ConsoleAuthor) AvatarURLWithSize(size int) string {
	return util.ImageSize(u.AvatarURL, size, size)
}

// ConsoleCategory represents console category.
type ConsoleCategory struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	Tags        string `json:"tags"`
}

// ConsoleComment represents console comment.
type ConsoleComment struct {
	ID            uint64         `json:"id"`
	Author        *ConsoleAuthor `json:"author"`
	ArticleAuthor *ConsoleAuthor `json:"articleAuthor"`
	CreatedAt     string         `json:"createdAt"`
	Title         string         `json:"title"`
	Content       template.HTML  `json:"content"`
	URL           string         `json:"url"`
}

// ConsoleNavigation represents console navigation.
type ConsoleNavigation struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	IconURL    string `json:"iconURL"`
	OpenMethod string `json:"openMethod"`
	Number     int    `json:"number"`
}

// ConsoleTheme represents console theme.
type ConsoleTheme struct {
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnailURL"`
}

// ConsoleUser represents console user.
type ConsoleUser struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	Role         int    `json:"role"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatarURL"`
	ArticleCount int    `json:"articleCount"`
}
