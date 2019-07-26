package models

import "github.com/satori/go.uuid"

type User struct {
	Model
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Password string `json:"_"`
	Salt     string `json:"_"`
	RoleUuid string `json:"roleUuid"`
	Role     Role   `json:"role" gorm:"ForeignKey:RoleUuid"`
}

func GetUsers(pageSize int, pageNo int, name string) (user []User, count int) {
	offset := (pageNo - 1) * pageSize
	db.Preload("Role").Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&user)
	db.Model(&User{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func EditUser(user *User, id string) bool {
	err := db.Model(&User{}).Where("uuid = ?", id).Updates(user).Error
	if err != nil {
		return false
	}
	return true
}

func AddUser(user *User) bool {
	err := db.Create(user).Error
	if err != nil {
		return false
	}
	return true
}

func DeleteUser(id string) bool {
	err := db.Delete(&User{}, "uuid = ?", id).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

// 1 不存在用户名
func AuthUserName(name string) string {
	user := User{
		Name: name,
	}
	if db.Select("salt").Where(&user).First(&user).RecordNotFound() {
		return ""
	}
	return user.Salt
}

func AuthUser(name, password string) (b bool, user *User) {
	user = &User{
		Name:     name,
		Password: password,
	}
	res := db.Where(&user).First(&user)
	if res.RecordNotFound() {
		b = false
	} else {
		b = true
	}
	return
}

func GetUser(uuid uuid.UUID) (user *User) {
	user = &User{}
	user.Uuid = uuid
	db.First(user)
	return
}
