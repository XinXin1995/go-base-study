package v1

import (
	"github.com/satori/go.uuid"
)

type ModulesApis struct {
	ModuleUuid uuid.UUID `json:"apiUuid"`
	ApiUuid    uuid.UUID `json:"apiUuid"`
}
type RolesModules struct {
	RoleUuid   uuid.UUID `json:"roleUuid"`
	ModuleUuid uuid.UUID `json:"modulesUuid"`
}
