package v1

import "github.com/satori/go.uuid"

type RolesModules struct {
	RoleUuid   uuid.UUID `json:"roleUuid"`
	ModuleUuid uuid.UUID `json:"modulesUuid"`
}
