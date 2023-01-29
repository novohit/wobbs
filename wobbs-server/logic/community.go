package logic

import (
	"wobbs-server/config"
	"wobbs-server/model"
)

func GetCommunityCategory() []model.Category {
	db := config.GetDB()
	category := make([]model.Category, 0)
	db.Where(&model.Category{}).Find(&category)
	//fmt.Println(users)
	return category
}
