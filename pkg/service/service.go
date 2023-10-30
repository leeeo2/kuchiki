package service

import "github.com/leeexeo/kuchiki/pkg/models"

var (
	userDao    = models.NewUserDao()
	articleDao = models.NewArticleDao()
)
