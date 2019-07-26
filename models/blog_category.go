package models

type Category struct {
	Model
	Name string `json:"name"`
}

func GetCategories(pageSize int, pageNo int, name string) (categories []Category, count int) {
	offset := (pageNo - 1) * pageSize
	db.Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&categories)
	db.Model(&Tag{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
	return
}

func AddCategory(category *Category) bool {
	err := db.Create(category).Error
	if err != nil {
		return false
	}
	return true
}

func EditCategory(category *Category, id string) bool {
	err := db.Model(&Category{}).Where("uuid = ?", id).Updates(category).Error
	if err != nil {
		return false
	}
	return true
}

func DelCategory(id string) bool {
	err := db.Delete(&Category{}, "uuid = ?", id).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
