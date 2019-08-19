package models

import (
	"github.com/satori/go.uuid"
)

type ArticleTags struct {
	ArticleUuid uuid.UUID `json:"articleUuid"`
	TagUuid     uuid.UUID `json:"tagUuid"`
}
