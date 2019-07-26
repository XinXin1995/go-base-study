package models

import (
	"github.com/satori/go.uuid"
)

type ModulesApis struct {
	ModuleUuid uuid.UUID `json:"apiUuid"`
	ApiUuid    uuid.UUID `json:"apiUuid"`
}
