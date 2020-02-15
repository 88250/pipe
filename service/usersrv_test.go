// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package service

import "testing"

func TestGetUserByName(t *testing.T) {
	user := User.GetUserByName(testPlatformAdminName)
	if nil == user {
		t.Errorf("user is nil")

		return
	}
	user = User.GetUserByName("notfound")
	if nil != user {
		t.Errorf("user should be nil")
	}

}

func TestGetUser(t *testing.T) {
	user := User.GetUser(uint64(1))
	if nil == user {
		t.Errorf("user is nil")

		return
	}
	if 1 != user.ID {
		t.Errorf("id is not [1]")
	}
}

func TestGetBlogUsers(t *testing.T) {
	users, _ := User.GetBlogUsers(1, 1)
	if 1 > len(users) {
		t.Errorf("users is empty")

		return
	}
}

func TestGetUserBlogs(t *testing.T) {
	blogs := User.GetUserBlogs(1)
	if 1 > len(blogs) {
		t.Errorf("blogs is tempty")

		return
	}
}
