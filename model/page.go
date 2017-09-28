// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
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

// Page types.
const (
	PageTypePage = iota
	PageTypeLink
)

// Page (customized navigation) model.
type Page struct {
	Model

	Title        string `gorm:"size:128" json:"title"`
	Content      string `gorm:"type:text" json:"content"`
	Permalink    string `gorm:"size:255" json:"permalink"`
	IconURL      string `gorm:"size:255" json:"iconURL"`
	Number       int    `json:"number"` // for sorting
	Type         int    `json:"type"`   // 0: page, 1: link
	Commentable  bool   `json:"commentable"`
	ViewCount    int    `json:"viewCount"`
	CommentCount int    `json:"commentCount"`

	BlogID uint
}
