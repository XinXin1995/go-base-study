package v1

import (
	"github.com/satori/go.uuid"
	"log"
)

type Role struct {
	Model
	Name    string   `json:"name"`
	Modules []Module `gorm:"many2many:role_modules;"`
}

func GetRoles(pageSize int, pageNo int, name string) (role []Role, count int) {
	offset := (pageNo - 1) * pageSize
	db.Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&role)
	db.Model(&Role{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func AddRole(role *Role) bool {
	err := db.Create(role).Error
	if err != nil {
		return false
	}
	return true
}

func EditRole(role *Role, id string) bool {
	err := db.Model(&Role{}).Where("uuid = ?", id).Updates(role).Error
	if err != nil {
		return false
	}
	return true
}

func DeleteRole(id string) bool {
	err := db.Delete(&Role{}, "uuid = ?", id).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetAllRoles() (roles []Role) {
	err := db.Find(&roles).Error
	if err != nil {
		log.Fatalln("mysql err: ", err)
	}
	return
}

func GetRoleModules(id string) (modules []Module) {
	role := &Role{}
	role.Uuid = uuid.FromStringOrNil(id)
	db.Model(role).Preload("ModuleGroup").Related(&modules, "Modules")
	return
}
