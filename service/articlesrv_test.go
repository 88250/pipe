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
	"testing"

	"github.com/b3log/solo.go/model"
)

func TestAddArticle(t *testing.T) {
	ConnectDB()

	article := &model.Article{AuthorID: 1,
		Title:       "Test 文章",
		Abstract:    "Test 摘要",
		Tags:        "Tag1, 标签2",
		Content:     "正文部分",
		Permalink:   "/test1",
		Status:      model.ArticleStatusPublished,
		Topped:      false,
		Commentable: true,
		Password:    "",
		ViewCount:   0,
	}

	Article.AddArticle(article)

}
