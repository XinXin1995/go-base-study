package v1

type Role struct {
	Model
	Name string `json:"name"`
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
