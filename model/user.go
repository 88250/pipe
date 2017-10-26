// Pipe - A small and beautiful blogging platform written in golang.
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

// User model.
type User struct {
	Model

	Name                  string `gorm:"size:32" json:"name"`
	Nickname              string `gorm:"size:32" json:"nickname"`
	Email                 string `gorm:"size:64" json:"email"`
	Role                  int    `json:"role"`
	AvatarURL             string `gorm:"size:255" json:"avatarURL"`
	B3Key                 string `gorm:"size:32" json:"b3Key"`
	Locale                string `gorm:"size:32 json:"locale"`
	ArticleCount          int    `json:"articleCount"`          // including drafts and published articles
	PublishedArticleCount int    `json:"publishedArticleCount"` // just including published articles
	Status                int    `json:"status"`

	BlogID uint `json:"blogID"`
}

// User roles.
const (
	UserRolePlatformAdmin = iota
	UserRoleBlogAdmin
	UserRoleBlogUser
	UserRoleBlogVisitor
)

// User status.
const (
	UserStatusOK = iota
	UserStatusLoginRestricted
)
