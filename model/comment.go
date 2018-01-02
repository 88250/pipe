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

import "time"

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

	BlogID uint64 `sql:"index" json:"blogID"`
}
