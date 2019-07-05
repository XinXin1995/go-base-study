package setting

import "github.com/jinzhu/gorm"

func CreteTable(name string, st interface{}, db *gorm.DB) {
	if !db.HasTable(TablePrefix + name) {
		db.CreateTable(st)
	}
}
