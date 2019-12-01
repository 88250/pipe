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
	"strconv"
	"testing"

	"github.com/88250/pipe/model"
)

const (
	articleRecordSize = 99
)

func TestGetArticleByPath(t *testing.T) {
	article := Article.GetArticleByPath("/hello-world", 1)
	if nil == article {
		t.Errorf("article is nil")
	}
}

var lastArticleID uint64

func TestAddArticle(t *testing.T) {
	for i := 0; i < articleRecordSize; i++ {
		article := &model.Article{AuthorID: 1,
			Title:       "Test 文章" + strconv.Itoa(i),
			Tags:        "Tag1, 标签2",
			Content:     "正文部分",
			Topped:      false,
			Commentable: true,
			BlogID:      1,
		}

		if err := Article.AddArticle(article); nil != err {
			t.Error("add article failed: " + err.Error())
		}

		lastArticleID = article.ID
	}

	statisticSetting := Statistic.GetStatistic(model.SettingNameStatisticArticleCount, 1)
	if "100" != statisticSetting.Value {
		t.Errorf("expected is [%s], actual is [%s]", "100", statisticSetting.Value)
	}
}

func TestGetPreviousNextArticle(t *testing.T) {
	article := Article.GetPreviousArticle(lastArticleID, 1)
	if nil == article {
		t.Errorf("article is nil")

		return
	}

	article = Article.GetNextArticle(article.ID, 1)
	if nil == article {
		t.Errorf("article is nil")

		return
	}

	if article.ID != lastArticleID {
		t.Errorf("it is not the next article")
	}
}

func TestConsoleGetArticles(t *testing.T) {
	articles, pagination := Article.ConsoleGetArticles("", 1, 1)
	if adminConsoleArticleListPageSize != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", adminConsoleArticleListPageSize, len(articles))
	}

	if articleRecordSize+1 /* including "Hello,World!" */ != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", articleRecordSize+1, pagination.RecordCount)
	}
}

func TestGetArticles(t *testing.T) {
	articles, pagination := Article.GetArticles("", 1, 1)
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

func TestConsoleGetArticle(t *testing.T) {
	article := Article.ConsoleGetArticle(lastArticleID)
	if nil == article {
		t.Errorf("article is nil")

		return
	}
	if lastArticleID != article.ID {
		t.Errorf("id is not [" + strconv.Itoa(int(lastArticleID)) + "]")
	}
}

func TestUpdateArticle(t *testing.T) {
	updatedTitle := "Updated title"
	article := Article.ConsoleGetArticle(lastArticleID)
	article.Title = updatedTitle
	if err := Article.UpdateArticle(article); nil != err {
		t.Errorf("update article failed: " + err.Error())

		return
	}

	article = Article.ConsoleGetArticle(lastArticleID)
	if updatedTitle != article.Title {
		t.Errorf("expected is [%s], actual is [%s]", updatedTitle, article.Title)
	}
}

func TestIncArticleViewCount(t *testing.T) {
	article := Article.ConsoleGetArticle(lastArticleID)
	oldCnt := article.ViewCount
	if err := Article.IncArticleViewCount(article); nil != err {
		t.Errorf("inc article view count failed: " + err.Error())

		return
	}

	if oldCnt+1 != article.ViewCount {
		t.Errorf("expected is [%d], actual is [%d]", oldCnt+1, article.ViewCount)
	}
}

func TestNormalizeTagStr(t *testing.T) {
	tagStr := normalizeTagStr("带 空 格1,分号2；顿号3、正常4")
	if "带空格1,分号2,顿号3,正常4" != tagStr {
		t.Errorf("exptected is [%s], actual is [%s]", "带空格1,分号2,顿号3,正常4", tagStr)
	}

	tagStr = normalizeTagStr("")
	if "待分类" != tagStr {
		t.Errorf("exptected is [%s], actual is [%s]", "待分类", tagStr)
	}
}

func TestTagArticle(t *testing.T) {
	article := Article.ConsoleGetArticle(lastArticleID)

	tx := db.Begin()
	if err := tagArticle(tx, article); nil != err {
		tx.Rollback()
		t.Errorf("tag article failed: " + err.Error())

		return
	}
	tx.Commit()
}

func TestRemoveArticle(t *testing.T) {
	article := Article.ConsoleGetArticle(lastArticleID)
	if nil == article {
		t.FailNow()

		return
	}

	if err := Article.RemoveArticle(lastArticleID, 1); nil != err {
		t.Error(err)
	}

	article = Article.ConsoleGetArticle(lastArticleID)
	if nil != article {
		t.Error("remove article failed")
	}
}

func TestNormalizeArticlePath(t *testing.T) {
	article := &model.Article{
		Path: "/aaa",
	}
	article.ID = 1

	if err := normalizeArticlePath(article); nil != err {
		t.Error(err)

		return
	}
	if "/aaa" != article.Path {
		t.Errorf("expected is [%s], actual is [%s]", "/aaa", article.Path)
	}

	article.Path = ""
	if err := normalizeArticlePath(article); nil != err {
		t.Error(err)

		return
	}
	if len("/articles/2017/11/02/1") != len(article.Path) {
		t.Errorf(article.Path)
	}
}
