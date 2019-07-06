package v1

import "log"

type ModuleGroup struct {
	Model
	Name   string
	Router string
}

func AddModuleGroup(mg *ModuleGroup) bool {
	err := db.Create(mg).Error
	if err != nil {
		log.Fatalln("mysql err: ", err)
		return false
	} else {
		return true
	}

}

func GetAllModules() (mgs []ModuleGroup) {
	db.Find(&mgs)
	return
}
