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

import (
	"testing"

	"github.com/88250/pipe/model"
)

func TestConsoleGetComments(t *testing.T) {
	comments, pagination := Comment.ConsoleGetComments("", 1, 1)
	if 1 != len(comments) {
		t.Errorf("expected is [%d], actual is [%d]", 1, len(comments))

		return
	}
	if 1 != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", 1, pagination.RecordCount)
	}
}

func TestGetRecentComments(t *testing.T) {
	comments := Comment.GetRecentComments(10, 1)
	if 1 != len(comments) {
		t.Errorf("expected is [%d], actual is [%d]", 1, len(comments))

		return
	}
}

func TestGetArticleComments(t *testing.T) {
	comments, pagination := Comment.GetArticleComments(1, 1, 1)
	if 0 != len(comments) {
		t.Errorf("expected is [%d], actual is [%d]", 0, len(comments))

		return
	}
	if 0 != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", 0, pagination.RecordCount)
	}
}

func TestRemoveComment(t *testing.T) {
	articles, _ := Article.GetArticles("", 1, 1)
	comment := &model.Comment{
		ArticleID: articles[0].ID,
		AuthorID:  1,
		Content:   "写博客需要坚持，相信积累后必然会有收获，我们一起努力加油 :smile:",
		BlogID:    1,
	}
	if err := Comment.AddComment(comment); nil != err {
		t.Errorf("add comment failed: " + err.Error())

		return
	}

	if err := Comment.RemoveComment(comment.ID, 1); nil != err {
		t.Error(err)

		return
	}

	comments, _ := Comment.ConsoleGetComments("", 1, 1)
	if 1 != len(comments) {
		t.Error("remove comment failed")
	}
}
