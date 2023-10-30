package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/leeexeo/kon/id"
	"github.com/leeexeo/kon/log"
	"github.com/leeexeo/kuchiki/pkg/domain"
	"github.com/leeexeo/kuchiki/pkg/models"
)

func AddArticle(ctx *gin.Context, input *domain.AddArticleInput) (*domain.AddArticleOutput, error) {
	articleId := id.NewWithPrefix("a-")
	a := &models.Article{
		Id:       articleId,
		Author:   input.Author,
		Markdown: input.Markdown,
		Path:     input.Path,
	}

	if err := articleDao.Save(a); err != nil {
		log.Error(ctx, "save article failed")
		return nil, errors.New("save article failed")
	} else {
		return &domain.AddArticleOutput{ArticleId: articleId}, nil
	}
}
