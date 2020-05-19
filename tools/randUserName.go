package tools

import (
	"math/rand"
	"time"
)

func RandUserName(lenght int) (string) {
	listObj := []byte ("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	name := make([]byte, lenght)

	rand.Seed(time.Now().Unix())// 随机种子 增加随机性
	for i := range name {
		name[i] = listObj[rand.Intn(len(listObj))]
	}
	return string(name)
}