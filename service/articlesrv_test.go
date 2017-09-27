// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
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

	"github.com/b3log/solo.go/model"
)

const (
	articleRecordSize = 100
)

func TestAddArticle(t *testing.T) {
	ConnectDB()

	for i := 0; i < articleRecordSize; i++ {
		article := &model.Article{AuthorID: 1,
			Title:       "Test 文章" + strconv.Itoa(i),
			Abstract:    "Test 摘要",
			Tags:        "Tag1, 标签2",
			Content:     "正文部分",
			Permalink:   "/test" + strconv.Itoa(i),
			Status:      model.ArticleStatusPublished,
			Topped:      false,
			Commentable: true,
		}

		Article.AddArticle(article)
		//time.Sleep(500 * time.Millisecond)
	}
}

func TestGetConsoleArticles(t *testing.T) {
	articles, pagination := Article.ConsoleGetArticles(1)

	if adminConsoleArticleListPageSize != len(articles) {
		t.Errorf("expected is [%d], actual is [%d]", adminConsoleArticleListPageSize, len(articles))
	}

	if articleRecordSize != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", articleRecordSize, pagination.RecordCount)
	}
}

func TestGetConsoleArticle(t *testing.T) {
	article := Article.ConsoleGetArticle(1)
	if nil == article {
		t.Errorf("article is nil")

		return
	}

	if 1 != article.ID {
		t.Errorf("id is not [1]")
	}
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
