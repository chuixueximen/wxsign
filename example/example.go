package main

import (
	"fmt"

	"github.com/chuixueximen/wxsign"
	redis "gopkg.in/redis.v3"
)

func init() {
	// 初始化缓存access_token及ticket的redis
	rdsClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	wxsign.WxSignRdsInit(rdsClient)
}

func main() {
	ws := wxsign.New(
		"wx863db1421192d7ea",
		"b3e8334fb24c39d3f4d1f9c076d82bcc",
		// 缓存access_token使用的redis key
		"wechat_access_token_redis_key",
		// 缓存ticket使用的redis key
		"wechat_jsapi_ticket_redis_key",
	)
	sign, err := ws.GetJsSign("https://api.kuailife456.com")
	if err != nil {
		fmt.Print("Get js sign err-> %#v", err)
		return
	}
	fmt.Print("Js Sign: %#v", sign)
}
