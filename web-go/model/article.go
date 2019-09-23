package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Text string
}

func GetAllArticles(db *gorm.DB) ([]Article, error) {
	articles := make([]Article, 10)
	if err := db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticleByID(db *gorm.DB, id uint) (*Article, error) {
	article := new(Article)
	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func PostArticle(db *gorm.DB, article *Article) error {
	return db.Create(&article).Error
}

func DeleteArticleByID(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Delete(Article{}).Error
}
