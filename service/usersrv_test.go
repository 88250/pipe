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
