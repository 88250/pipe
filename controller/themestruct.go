package controller

type ThemeListArticle struct {
	ID           uint
	Abstract     string
	Author       *ThemeAuthor
	CreatedAt    string
	Title        string
	Tags         []*ThemeTag
	URL          string
	Topped       bool
	ViewCount    int
	CommentCount int
	ThumbnailURL string
	Content      string
}

type ThemeTag struct {
	Title string
	URL   string
}

type ThemeAuthor struct {
	Name      string
	AvatarURL string
	URL       string
}

type ThemeCategory struct {
	Title string
	URL   string
}

type ThemeListComment struct {
	ID        uint
	Title     string
	Content   string
	URL       string
	Author    *ThemeAuthor
	CreatedAt string
}
