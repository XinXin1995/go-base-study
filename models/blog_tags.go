package models

type Tag struct {
	Model
	Name   string `json:"name"`
	Hit    int    `json:"hit"`
	Effect string `json:"effect"`
	Color  string `json:"color"`
}

func GetAllTags(name string) (tags []Tag) {
	db.Where("name LIKE ?", "%"+name+"%").Find(&tags)
	return
}

func GetTags(pageSize int, pageNo int, name string) (tags []Tag, count int) {
	offset := (pageNo - 1) * pageSize
	db.Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&tags)
	db.Model(&Tag{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func AddTag(tag *Tag) bool {
	err := db.Create(tag).Error
	if err != nil {
		return false
	}
	return true
}

func EditTag(tag *Tag, id string) bool {
	err := db.Model(&Tag{}).Where("uuid = ?", id).Updates(tag).Error
	if err != nil {
		return false
	}
	return true
}

func DelTag(id string) bool {
	err := db.Delete(&Tag{}, "uuid = ?", id).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
func GetTagsById(tagIds []string) (tags []Tag) {
	db.Where("uuid in (?)", tagIds).Find(&tags)
	return
}
