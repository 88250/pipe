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
	"github.com/88250/pipe/util"
)

// User model.
type User struct {
	Model

	Name              string `gorm:"size:32" json:"name"`
	Nickname          string `gorm:"size:32" json:"nickname"`
	AvatarURL         string `gorm:"size:255" json:"avatarURL"`
	B3Key             string `gorm:"size:32" json:"b3Key"`
	Locale            string `gorm:"size:32" json:"locale"`
	TotalArticleCount int    `json:"totalArticleCount"`
	GithubId          string `gorm:"255" json:"githubId"`
}

// User roles.
const (
	UserRoleNoLogin = iota
	UserRolePlatformAdmin
	UserRoleBlogAdmin
	UserRoleBlogUser
)

// AvatarURLWithSize returns avatar URL with the specified size.
func (u *User) AvatarURLWithSize(size int) string {
	return util.ImageSize(u.AvatarURL, size, size)
}
