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
