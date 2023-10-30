package domain

type AddArticleInput struct {
	Author   string
	Markdown string
	Path     string
}

type AddArticleOutput struct {
	ArticleId string
}
