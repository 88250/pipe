// Solo.go - A small and beautiful blogging platform written in golang.
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

package service

import (
	"testing"
)

func TestConsoleGetComments(t *testing.T) {
	comments, pagination := Comment.ConsoleGetComments(1, 1)

	if 1 != len(comments) {
		t.Errorf("expected is [%d], actual is [%d]", 1, len(comments))
	}
	if 1 != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", 1, pagination.RecordCount)
	}
}

func TestRemoveComment(t *testing.T) {
	if err := Comment.RemoveComment(1); nil != err {
		t.Error(err)
	}

	comments, _ := Comment.ConsoleGetComments(1, 1)
	if 0 != len(comments) {
		t.Error("remove comment failed")
	}
}
