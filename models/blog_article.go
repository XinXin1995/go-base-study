package models

import "github.com/satori/go.uuid"

type Article struct {
	Model
	Name         string    `json:"name"`
	Content      string    `json:"name"gorm:"type:text"`
	Creator      uuid.UUID `json:"creator"`
	Tags         []Tag     `json:"tags"gorm:"many2many:articles_tags"`
	CategoryUuid string    `json:"categoryUuid"`
	Category     Category  `json:"category"`
}
