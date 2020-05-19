package tools

import (
	"demo_items/gin_project/gin_vue_v2/model"
	"github.com/jinzhu/gorm"
)

func IsUserNameExist(db *gorm.DB, uname string) (bool) {
	var user model.User
	db.Where("name=?", uname).Find(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
