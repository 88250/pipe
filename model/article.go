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
	"time"
)

// Article model.
type Article struct {
	Model

	AuthorID uint64 `json:"authorID" structs:"authorID"`
	Title    string `gorm:"size:128" json:"title" structs:"title"`
	//Abstract     string    `gorm:"type:mediumtext" json:"abstract" structs:"abstract"`
	Abstract string `gorm:"type:text" json:"abstract" structs:"abstract"`
	Tags     string `gorm:"type:text" json:"tags" structs:"tags"`
	//Content      string    `gorm:"type:mediumtext" json:"content" structs:"content"`
	Content      string    `gorm:"type:text" json:"content" structs:"content"`
	Path         string    `sql:"index" gorm:"size:255" json:"path" structs:"path"`
	Status       int       `sql:"index" json:"status" structs:"status"`
	Topped       bool      `json:"topped" structs:"topped"`
	Commentable  bool      `json:"commentable" structs:"commentable"`
	ViewCount    int       `json:"viewCount" structs:"viewCount"`
	CommentCount int       `json:"commentCount" structs:"commentCount"`
	IP           string    `gorm:"size:128" json:"ip" structs:"ip"`
	UserAgent    string    `gorm:"size:255" json:"userAgent" structs:"userAgent"`
	PushedAt     time.Time `json:"pushedAt" structs:"pushedAt"`

	BlogID uint64 `sql:"index" json:"blogID" structs:"blogID"`
}

// Article statuses.
const (
	ArticleStatusOK = iota
)
