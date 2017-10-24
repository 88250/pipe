package console

import "html/template"

type ConsoleArticle struct {
	ID           uint           `json:"id"`
	Author       *ConsoleAuthor `json:"author"`
	CreatedAt    string         `json:"createdAt"`
	Title        string         `json:"title"`
	Tags         []*ConsoleTag  `json:"tags"`
	URL          string         `json:"url"`
	Topped       bool           `json:"topped"`
	ViewCount    int            `json:"viewCount"`
	CommentCount int            `json:"commentCount"`
}

type ConsoleTag struct {
	Title string `json:"title"`
	URL   string `json:"url,omitempty"`
}

type ConsoleAuthor struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarURL"`
}

type ConsoleCategory struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	Tags        string `json:"tags"`
}

type ConsoleComment struct {
	ID            uint           `json:"id"`
	Author        *ConsoleAuthor `json:"author"`
	ArticleAuthor *ConsoleAuthor `json:"articleAuthor"`
	CreatedAt     string         `json:"createdAt"`
	Title         string         `json:"title"`
	Content       template.HTML  `json:"content"`
	URL           string         `json:"url"`
}

type ConsoleNavigation struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	IconURL    string `json:"iconURL"`
	OpenMethod string `json:"openMethod"`
	Number     int    `json:"number"`
}

type ConsoleTheme struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	PreviewURL   string `json:"previewURL"`
	ThumbnailURL string `json:"thumbnailURL"`
}
