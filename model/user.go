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
	"strconv"
)

// User model.
type User struct {
	Model

	Name              string `gorm:"size:32" json:"name"`
	Nickname          string `gorm:"size:32" json:"nickname"`
	AvatarURL         string `gorm:"size:255" json:"avatarURL"`
	B3Key             string `gorm:"size:32" json:"b3Key"`
	Locale            string `gorm:"size:32 json:"locale"`
	TotalArticleCount int    `json:"totalArticleCount"`
}

// User roles.
const (
	UserRoleNoLogin = iota
	UserRolePlatformAdmin
	UserRoleBlogAdmin
	UserRoleBlogUser
)

func (u *User) AvatarURLWithSize(size int) string {
	return u.AvatarURL + "?imageView2/1/w/" + strconv.Itoa(size) + "/h/" + strconv.Itoa(size) + "/interlace/1/q/100"
}
