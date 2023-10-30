package models

type Article struct {
	Id       string `json:"Id" gorm:"size:64;not null;primaryKey;comment:文章ID"`
	Author   string `json:"Author" gorm:"size:64;comment:作者用户ID"`
	Markdown string `json:"Markdown" gorm:"type:longtext;comment:markdown文本"`
	Path     string `json:"Path" gorm:"size:256;comment:本地文本路径"`
	ModelTime
}

func (a *Article) TableName() string {
	return "articles"
}

type ArticleDao interface {
	Save(*Article) error
	Delete(id string) error
	Update(*Article) error
	DescribeArticle(id string) (*Article, error)
	DescribeArticles(input *DescribeInput) ([]*Article, error)
}

func NewArticleDao() ArticleDao {
	return &ArticleDaoImpl{}
}

type ArticleDaoImpl struct{}

func (a *ArticleDaoImpl) Save(article *Article) error {
	return GetDb().Create(article).Error
}

func (a *ArticleDaoImpl) Delete(id string) error {
	return GetDb().Where("id = ?", id).Delete(&Article{}).Error
}

func (a *ArticleDaoImpl) Update(article *Article) error {
	tx := GetDb().Model(&Article{})
	tx.Where("id = ?", article.Id)
	fields := map[string]interface{}{
		"author":   article.Author,
		"markdown": article.Markdown,
		"path":     article.Path,
	}
	return tx.Updates(fields).Error
}

func (a *ArticleDaoImpl) DescribeArticle(id string) (*Article, error) {
	tx := GetDb().Where("id = ?", id)
	var article Article
	if err := tx.First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *ArticleDaoImpl) DescribeArticles(input *DescribeInput) ([]*Article, error) {
	tx := GetDb().Where(input.Query, input.Params...)
	if input.PageSize > 0 && input.PageNumber > 0 {
		tx = tx.Limit(input.PageSize).Offset((input.PageNumber - 1) * input.PageSize)
	}
	if len(input.Order) > 0 {
		tx.Order(input.Order)
	}

	articles := make([]*Article, 0)
	if err := tx.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
