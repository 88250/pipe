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
