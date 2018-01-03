// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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

package model

import (
	"html/template"
	"strconv"
)

type ThemeArticle struct {
	ID           uint64        `json:",omitempty"`
	Abstract     template.HTML `json:"abstract"`
	Author       *ThemeAuthor  `json:",omitempty"`
	CreatedAt    string        `json:",omitempty"`
	Title        string        `json:"title"`
	Tags         []*ThemeTag   `json:"tags"`
	URL          string        `json:"url"`
	Topped       bool          `json:",omitempty"`
	ViewCount    int           `json:",omitempty"`
	CommentCount int           `json:",omitempty"`
	ThumbnailURL string        `json:",omitempty"`
	Content      template.HTML `json:",omitempty"`
	Editable     bool          `json:",omitempty"`
}

type ThemeTag struct {
	Title        string `json:"title"`
	URL          string `json:"url"`
	ArticleCount int    `json:",omitempty"`
}

type ThemeArchive struct {
	Title        string
	URL          string
	ArticleCount int
}

type ThemeAuthor struct {
	Name         string
	AvatarURL    string
	URL          string
	ArticleCount int
}

func (author *ThemeAuthor) AvatarURLWithSize(size int) string {
	return author.AvatarURL + "?imageView2/1/w/" + strconv.Itoa(size) + "/h/" + strconv.Itoa(size) + "/interlace/1/q/100"
}

type ThemeCategory struct {
	Title        string
	URL          string
	Description  string
	Tags         []*ThemeTag
	ArticleCount int
}

type ThemeComment struct {
	ID         uint64
	Title      string
	Content    template.HTML
	URL        string
	Author     *ThemeAuthor
	CreatedAt  string
	Removable  bool
	ReplyCount int
	Parent     *ThemeComment
}

type ThemeReply struct {
	ID        uint64
	Content   template.HTML
	URL       string
	Author    *ThemeAuthor
	CreatedAt string
}
