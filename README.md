# wxsign
[![Build Status](https://travis-ci.org/usth/wxsign.svg?branch=master)](https://travis-ci.org/chuixueximen/wxsign)
[![Go Report Card](https://goreportcard.com/badge/github.com/chuixueximen/wxsign)](https://goreportcard.com/report/github.com/chuixueximen/wxsign)
[![GoDoc](http://godoc.org/github.com/chuixueximen/wxsign?status.svg)](http://godoc.org/github.com/chuixueximen//wxsign)

微信公众号二次分享、请求Js签名.

- 更多逻辑已经迁移到[owechat](https://github.com/chuixueximen/owechat)

## 功能
- 获取微信分享所需要的js签名信息

- 返回签名信息
```
{
    Appid     string `json:"appid"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}	
```

## 安装

```
go get github.com/chuixueximen/wxsign
```

## 使用

```
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
		"appid",
		"secret",
		// 缓存access_token使用的redis key
		"wxsign:token",
		// 缓存ticket使用的redis key
		"wxsign:ticket",
	)
	sign, err := ws.GetJsSign("https://www.ooz.ink")
	if err != nil {
		fmt.Print("Get js sign err-> %#v", err)
		return
	}
	fmt.Print("Js Sign: %#v", sign)
}
```