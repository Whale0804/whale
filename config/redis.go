package config

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var Cache cache.Cache

func InitRedis() {
	redis_host := beego.AppConfig.String("redis_host")
	redis_port := beego.AppConfig.String("redis_port")
	redis_pass := beego.AppConfig.String("redis_pass")
	redis_db := beego.AppConfig.String("redis_db")
	redisConf := fmt.Sprintf(`{"key":"%s","conn":"%s:%s","dbNum":"%d","password":"%s"}`,
		"admin",
		redis_host,
		redis_port,
		redis_db,
		redis_pass,
	)
	var err error
	Cache, err = cache.NewCache("redis", redisConf)
	if err != nil {
		fmt.Println(redisConf)
		fmt.Println("Redis connection fail:" + err.Error())
	}

}
