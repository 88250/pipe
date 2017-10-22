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

package service

import (
	"strconv"
	"testing"

	"github.com/b3log/pipe/model"
)

const (
	articleRecordSize = 99
)

func TestGetArticleByPath(t *testing.T) {
	article := Article.GetArticleByPath("/hello-world")
	if nil == article {
		t.Errorf("article is nil")
	}
}

func TestAddArticle(t *testing.T) {
	for i := 0; i < articleRecordSize; i++ {
		article := &model.Article{AuthorID: 1,
			Title:       "Test 文章" + strconv.Itoa(i),
			Tags:        "Tag1, 标签2",
			Content:     "正文部分",
			Path:        "/test" + strconv.Itoa(i),
			Status:      model.ArticleStatusPublished,
			Topped:      false,
			Commentable: true,
			BlogID:      1,
		}

		if err := Article.AddArticle(article); nil != err {
			t.Error("add article failed: " + err.Error())
		}
	}
}

func TestGetPreviousArticle(t *testing.T) {
	article := Article.GetPreviousArticle(2, 1)
	if nil == article {
		t.Errorf("article is nil")

		return
	}
	if article.ID >= 2 {
		t.Errorf("it is not the previous article")
	}
}

func TestGetNextArticle(t *testing.T) {
	article := Article.GetNextArticle(1, 1)
	if nil == article {
		t.Errorf("article is nil")

		return
	}
	if article.ID <= 1 {
		t.Errorf("it is not the next article")
	}
}

func TestConsoleGetArticles(t *testing.T) {
	articles, pagination := Article.ConsoleGetArticles(1, 1)
	if adminConsoleArticleListPageSize != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", adminConsoleArticleListPageSize, len(articles))
	}

	if articleRecordSize+1 /* including "Hello,World!" */ != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", articleRecordSize+1, pagination.RecordCount)
	}
}

func TestGetArticles(t *testing.T) {
	articles, pagination := Article.GetArticles(1, 1)
	if 20 != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", 20, len(articles))

		return
	}
	if articleRecordSize+1 /* including "Hello,World!" */ != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", articleRecordSize+1, pagination.RecordCount)
	}
}

func TestGetMostViewArticles(t *testing.T) {
	articles := Article.GetMostViewArticles(10, 1)
	if 10 != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", 10, len(articles))
	}
}

func TestGetMostCommentArticles(t *testing.T) {
	articles := Article.GetMostCommentArticles(10, 1)
	if 10 != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", 10, len(articles))
	}
}

func TestGetRandomArticles(t *testing.T) {
	articles := Article.GetRandomArticles(10, 1)
	if 10 != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", 10, len(articles))
	}
}

func TestConsoleGetArticle(t *testing.T) {
	article := Article.ConsoleGetArticle(1)
	if nil == article {
		t.Errorf("article is nil")

		return
	}
	if 1 != article.ID {
		t.Errorf("id is not [1]")
	}
}

func TestUpdateArticle(t *testing.T) {
	updatedTitle := "Updated title"
	article := Article.ConsoleGetArticle(1)
	article.Title = updatedTitle
	if err := Article.UpdateArticle(article); nil != err {
		t.Errorf("update article failed: " + err.Error())

		return
	}

	article = Article.ConsoleGetArticle(1)
	if updatedTitle != article.Title {
		t.Errorf("expected is [%s], actual is [%s]", updatedTitle, article.Title)
	}
}

func TestNormalizeTagStr(t *testing.T) {
	tagStr := normalizeTagStr("带 空 格1,分号2；顿号3、正常4")
	if "带空格1,分号2,顿号3,正常4" != tagStr {
		t.Error("exptected is [%s], actual is [%s]", "带空格1,分号2,顿号3,正常4", tagStr)
	}
}

func TestTagArticle(t *testing.T) {
	article := Article.ConsoleGetArticle(1)

	tx := db.Begin()
	if err := tagArticle(tx, article); nil != err {
		tx.Rollback()
		t.Errorf("tag article failed: " + err.Error())

		return
	}
	tx.Commit()
}

func TestRemoveArticle(t *testing.T) {
	if err := Article.RemoveArticle(1); nil != err {
		t.Error(err)
	}

	article := Article.ConsoleGetArticle(1)
	if nil != article {
		t.Error("remove article failed")
	}
}
