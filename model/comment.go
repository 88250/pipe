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
	"math"
	"time"
)

// Comment model.
type Comment struct {
	Model

	ArticleID       uint64    `json:"articleID"`
	AuthorID        uint64    `json:"authorID"`
	Content         string    `gorm:"type:text" json:"content"`
	ParentCommentID uint64    `json:"parentCommentID"` // ID of replied comment
	IP              string    `gorm:"size:128" json:"ip"`
	UserAgent       string    `gorm:"size:255" json:"userAgent"`
	PushedAt        time.Time `json:"pushedAt"`

	AuthorName      string `gorm:"size:32" json:"authorName"`       // exist if this comment sync from Sym, https://github.com/b3log/pipe/issues/98
	AuthorAvatarURL string `gorm:"size:255" json:"authorAvatarURL"` // exist if this comment sync from Sym, https://github.com/b3log/pipe/issues/98
	AuthorURL       string `gorm:"size:255" json:"authorURL"`       // exist if this comment sync from Sym, https://github.com/b3log/pipe/issues/98

	BlogID uint64 `sql:"index" json:"blogID"`
}

// SyncCommentAuthorID is the id of sync comment bot.
const SyncCommentAuthorID = math.MaxInt32
