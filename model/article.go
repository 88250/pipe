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
	"time"
)

// Article model.
type Article struct {
	Model

	AuthorID     uint64    `json:"authorID"`
	Title        string    `gorm:"size:128" json:"title"`
	Tags         string    `gorm:"type:text" json:"tags"`
	Content      string    `gorm:"type:mediumtext" json:"content"`
	Path         string    `sql:"index" gorm:"size:255" json:"path"`
	Status       int       `sql:"index" json:"status"`
	Topped       bool      `json:"topped"`
	Commentable  bool      `json:"commentable"`
	ViewCount    int       `json:"viewCount"`
	CommentCount int       `json:"commentCount"`
	IP           string    `gorm:"size:128" json:"ip"`
	UserAgent    string    `gorm:"size:255" json:"userAgent"`
	PushedAt     time.Time `json:"pushedAt"`

	BlogID uint64 `sql:"index" json:"blogID"`
}

// Article statuses.
const (
	ArticleStatusOK = iota
)
