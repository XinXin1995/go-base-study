package v1

import "log"

type Api struct {
	Model
	Name string `json:"name"`
	Path string `json:"path"`
}

func AddApi(api *Api) bool {
	err := db.Create(api).Error
	if err != nil {
		log.Fatalln("mysql err: ", err)
		return false
	} else {
		return true
	}
}

func GetApis(pageSize int, pageNo int, name string) (api []Api, count int) {
	offset := (pageNo - 1) * pageSize
	db.Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&api)
	db.Model(&Api{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func DeleteApi(id string) bool {
	db := db.Delete(&Api{}, "uuid = ?", id)
	if db.Error != nil {
		return false
	} else {
		return true
	}
}

//0：操作成功 1：没有这条数据 2： 操作失败
func EditApi(api *Api, id string) bool {
	db := db.Model(&Api{}).Where("uuid = ?", id).Updates(api)
	if db.Error != nil {
		return false
	} else {
		return true
	}
}
