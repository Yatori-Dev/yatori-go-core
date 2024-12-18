package xuexitong

import (
	"github.com/yatori-dev/yatori-go-core/api/xuexitong"
	log2 "github.com/yatori-dev/yatori-go-core/utils/log"
)

func XueXiTLoginAction(cache *xuexitong.XueXiTUserCache) error {
	_, err := cache.LoginApi()
	if err == nil {
		log2.Print(log2.INFO, "["+cache.Name+"] "+" 登录成功")
	}
	return nil
}
