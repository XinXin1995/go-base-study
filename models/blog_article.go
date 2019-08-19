package models

import "github.com/satori/go.uuid"

type Article struct {
	Model
	Name         string    `json:"name"`
	Content      string    `json:"name" gorm:"type:text"`
	Creator      uuid.UUID `json:"creator"`
	Tags         []Tag     `json:"tags" gorm:"many2many:article_tags"`
	CategoryUuid uuid.UUID `json:"categoryUuid"`
	Category     Category  `json:"category"`
	IsDraft      int       `json:"isDraft"`
}
type ArticleParam struct {
	Name         string   `json:"name"`
	Content      string   `json:"content" `
	Creator      string   `json:"creator"`
	Tags         []string `json:"tags"`
	CategoryUuid string   `json:"categoryUuid"`
	IsDraft      int      `json:"isDraft"`
}

func GetArticles(pageSize int, pageNo int, name string) (articles []Article, count int) {
	offset := (pageNo - 1) * pageSize
	db.Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&articles)
	db.Model(&Article{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func AddArticle(article *Article) bool {
	err := db.Create(article).Error
	if err != nil {
		return false
	}
	return true
}

func EditArticle(article *Article) bool {
	err := db.Model(&Article{}).Updates(article).Error
	if err != nil {
		return false
	}
	return true
}

func DelArticle(id string) bool {
	err := db.Delete(&Article{}, "uuid = ?", id).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
