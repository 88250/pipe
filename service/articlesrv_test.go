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

	ArticleService.AddArticle(article)

}
