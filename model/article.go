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

package model

import (
	"time"
)

// Article model.
type Article struct {
	Model

	AuthorID     uint64    `json:"authorID" structs:"authorID"`
	Title        string    `gorm:"size:128" json:"title" structs:"title"`
	Abstract     string    `gorm:"type:mediumtext" json:"abstract" structs:"abstract"`
	Tags         string    `gorm:"type:text" json:"tags" structs:"tags"`
	Content      string    `gorm:"type:mediumtext" json:"content" structs:"content"`
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
