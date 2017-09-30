package console

import (
	"net/http"

	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

type ConsoleComment struct {
	ID            uint    `json:"id"`
	Author        *Author `json:"author"`
	ArticleAuthor *Author `json:"articleAuthor"`
	CreatedAt     string  `json:"createdAt"`
	Title         string  `gorm:"size:128" json:"title"`
	Content       string  `gorm:"type:text" json:"content"`
	Permalink     string  `json:"permalink"`
}

func GetCommentsCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)
	commentModels, pagination := service.Comment.ConsoleGetComments(c.GetInt("p"), sessionData.BID)

	comments := []*ConsoleComment{}
	for _, commentModel := range commentModels {
		author := &Author{
			Name:      commentModel.AuthorName,
			AvatarURL: commentModel.AuthorAvatarURL,
		}
		articleAuthor := &Author{
			Name:      "article author name",
			AvatarURL: "article author avatar URL",
		}

		comment := &ConsoleComment{
			ID:            commentModel.ID,
			Author:        author,
			ArticleAuthor: articleAuthor,
			CreatedAt:     commentModel.CreatedAt.Format("2006-01-02"),
			Title:         "article title",
			Content:       commentModel.Content,
			Permalink:     sessionData.BPath + "todo comment link",
		}

		comments = append(comments, comment)
	}

	data := map[string]interface{}{}
	data["comments"] = comments
	data["pagination"] = pagination
	result.Data = data
}
