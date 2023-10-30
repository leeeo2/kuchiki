package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leeexeo/kon/log"
	"github.com/leeexeo/kuchiki/pkg/domain"
	"github.com/leeexeo/kuchiki/pkg/service"
)

func validateAddArticleInput(input *domain.AddArticleInput) error {
	return nil
}

func AddArticle(ctx *gin.Context) {
	var input domain.AddArticleInput
	if err := ctx.Bind(&input); err != nil {
		log.Error(ctx, "bind AddArticleInput failed", "error", err)
		Response(ctx, nil, err)
		return
		// ctx.JSON(http.StatusBadRequest, gin.H{"message": "bind AddArticleInput failed", "error": err.Error()})
	}
	if err := validateAddArticleInput(&input); err != nil {
		Response(ctx, nil, err)
		return
		// ctx.JSON(http.StatusBadRequest, gin.H{"message": "validate parameters failed", "error": err.Error()})
	}
	output, err := service.AddArticle(ctx, &input)
	Response(ctx, output, err)
}
