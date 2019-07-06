package v1

type Module struct {
	Model
	Name        string `json:"name"`
	Router      string `json:"router"`
	GroupUuid   string `json:"groupUuid" gorm:"ForeignKey:GroupUuid"`
	ModuleGroup ModuleGroup
	Apis        []Api `gorm:"many2many: modules_Apis"`
}
