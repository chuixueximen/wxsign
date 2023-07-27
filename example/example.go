package main

import (
	"fmt"
	"github.com/chuixueximen/wxsign"
	"github.com/go-redis/redis/v8"
)

func init() {
	// 初始化缓存access_token及ticket的redis
	rdsClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "",
		DB:       0,
	})
	wxsign.WxSignRdsInit(rdsClient)
}

func main() {
	ws := wxsign.New(
		"appid",
		"secret",
		// 缓存access_token使用的redis key
		"wechat_access_token_redis_key",
		// 缓存ticket使用的redis key
		"wechat_jsapi_ticket_redis_key",
	)
	sign, err := ws.GetJsSign("https://www.xxxx.com")
	if err != nil {
		fmt.Print("Get js sign err-> %#v", err)
		return
	}
	fmt.Print("Js Sign: %#v", sign)
}
