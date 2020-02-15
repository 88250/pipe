// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package model

import (
	"html/template"

	"github.com/88250/pipe/util"
)

// ThemeArticle represents theme article.
type ThemeArticle struct {
	ID             uint64        `json:",omitempty"`
	Abstract       template.HTML `json:"abstract"`
	Author         *ThemeAuthor  `json:",omitempty"`
	CreatedAt      string        `json:",omitempty"`
	CreatedAtYear  string        `json:",omitempty"`
	CreatedAtMonth string        `json:",omitempty"`
	CreatedAtDay   string        `json:",omitempty"`
	Title          string        `json:"title"`
	Tags           []*ThemeTag   `json:"tags"`
	URL            string        `json:"url"`
	Topped         bool          `json:",omitempty"`
	ViewCount      int           `json:",omitempty"`
	CommentCount   int           `json:",omitempty"`
	ThumbnailURL   string        `json:",omitempty"`
	Content        template.HTML `json:",omitempty"`
	Editable       bool          `json:",omitempty"`
}

// ThemeTag represents theme tag.
type ThemeTag struct {
	Title        string `json:"title"`
	URL          string `json:"url"`
	ArticleCount int    `json:",omitempty"`
}

// ThemeArchive represents theme archive.
type ThemeArchive struct {
	Title        string
	URL          string
	ArticleCount int
}

// ThemeAuthor represents theme author.
type ThemeAuthor struct {
	Name         string
	AvatarURL    string
	URL          string
	ArticleCount int
}

// AvatarURLWithSize returns avatar URL with the specified size.
func (author *ThemeAuthor) AvatarURLWithSize(size int) string {
	return util.ImageSize(author.AvatarURL, size, size)
}

// ThemeCategory represents theme category.
type ThemeCategory struct {
	Title        string
	URL          string
	Description  string
	Tags         []*ThemeTag
	ArticleCount int
}

// ThemeComment represents theme comment.
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

// ThemeReply represents theme reply.
type ThemeReply struct {
	ID        uint64
	Content   template.HTML
	URL       string
	Author    *ThemeAuthor
	CreatedAt string
}
