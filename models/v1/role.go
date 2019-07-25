package v1

import (
	"github.com/satori/go.uuid"
	"log"
)

type Role struct {
	Model
	Name    string   `json:"name"`
	Modules []Module `gorm:"many2many:roles_modules;"`
}

type RoleModules struct {
	Id      string   `json:"id"`
	Modules []string `json:"modules"`
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

func GetRoleModules(roleUuid uuid.UUID) (modules []Module) {
	role := &Role{}
	role.Uuid = roleUuid
	db.Model(role).Related(&modules, "Modules")
	return
}

func AddRoleModules(roleUuid uuid.UUID, modules []string) bool {
	tx := db.Begin()
	if err := tx.Delete(RolesModules{}, "role_uuid = ?", roleUuid).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}
	for _, v := range modules {
		moduleUuid, err := uuid.FromString(v)
		if err == nil {
			ra := &RolesModules{
				RoleUuid:   roleUuid,
				ModuleUuid: moduleUuid,
			}
			if err := tx.Create(ra).Error; err != nil {
				log.Println(err)
				tx.Rollback()
				return false
			}
		} else {
			tx.Rollback()
			return false
		}
	}
	tx.Commit()
	return true
}
