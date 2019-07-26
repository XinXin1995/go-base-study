package models

import (
	"github.com/satori/go.uuid"
	"log"
)

type Module struct {
	Model
	Name     string `json:"name" form:"name" binding:""`
	Router   string `json:"router"`
	Icon     string `json:"icon"`
	Apis     []Api  `json:"apis" gorm:"many2many:modules_apis"`
	ParentId string `json:"parentId"`
}

type ModuleApis struct {
	Id   string   `json:"id"`
	Apis []string `json:"apis"`
}

func AddModule(module *Module) bool {
	err := db.Create(module).Error
	if err != nil {
		log.Fatalln("mysql err: ", err)
		return false
	} else {
		return true
	}
}

func GetModules(pageSize int, pageNo int, name string) (module []Module, count int) {
	offset := (pageNo - 1) * pageSize
	db.Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&module)
	db.Model(&Module{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func DeleteModule(id string) int {
	db := db.Delete(&Module{}, "uuid = ?", id)
	if db.RowsAffected == 0 {
		return 1
	} else if db.Error != nil {
		return 2
	} else {
		return 0
	}
}

//0：操作成功 1：没有这条数据 2： 操作失败
func EditModule(module *Module, id string) bool {
	db := db.Model(&Module{}).Where("uuid = ?", id).Updates(module)
	if db.Error != nil {
		return false
	} else {
		return true
	}
}

func GetAllModules() (modules []Module) {
	db.Find(&modules)
	return
}

func AddModuleApis(apis []string, moduleUuid uuid.UUID) bool {
	tx := db.Begin()
	if err := tx.Delete(ModulesApis{}, "module_uuid = ?", moduleUuid).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}
	for _, v := range apis {
		apiUuid, err := uuid.FromString(v)
		if err == nil {
			ma := &ModulesApis{
				ModuleUuid: moduleUuid,
				ApiUuid:    apiUuid,
			}
			if err := tx.Create(ma).Error; err != nil {
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

func GetModuleApis(moduleUuid uuid.UUID) (apis []Api) {
	//"SELECT B.*  FROM `modules_apis` A LEFT JOIN `api` B ON B.uuid = A.api_uuid  WHERE A.module_uuid = ?", moduleUuid
	module := &Module{}
	module.Uuid = moduleUuid
	db.Model(module).Related(&apis, "Apis")
	return
}
