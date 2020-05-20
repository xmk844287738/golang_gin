package dto

// 去除转发前台的数据中，去除不想暴露在前台的数据
import (
	"demo_items/gin_project/gin_vue_v2/model"
)

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

//func GetUserInfoDto(value interface{}) (UserDto) {  // 第一种方法
//	// 空接口转结构体
//	user, exist := value.(model.User)
//	if !exist {
//		log.Printf("数据转换错误:%v\n", exist)
//	}
//
//	userdto := UserDto{
//		Name: user.Name,
//		Telephone: user.Telephone,
//	}
//	return userdto
//}

func GetUserInfoDto(user model.User) (UserDto) {
		userdto := UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
	return userdto
}
