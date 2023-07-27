package wxsign

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdsCli *redis.Client
	ctx    = context.Background()
)

// Init 初始化 redis client
func WxSignRdsInit(rc *redis.Client) {
	if rdsCli == nil {
		rdsCli = rc
	}
}

// PushToken 将微信token 存到 redis 中
func (wSign *WxSign) PushTokenByCache(token string, duration time.Duration) {
	rdsCli.Set(ctx, wSign.TokenRdsKey, token, duration)
}

// PushTicket 将微信jsticket 存到 redis 中
func (wSign *WxSign) PushTicketByCache(token string, duration time.Duration) {
	rdsCli.Set(ctx, wSign.TicketRdsKey, token, duration).Result()
}

// GetTokenByCache 从缓存获取access_token
func (wSign *WxSign) GetTokenByCache() string {
	var data string

	if val, err := rdsCli.Get(ctx, wSign.TokenRdsKey).Result(); err == nil {
		data = val
	}
	return data
}

// GetTicketByCache 从缓存获取ticket
func (wSign *WxSign) GetTicketByCache() string {
	var data string

	if val, err := rdsCli.Get(ctx, wSign.TicketRdsKey).Result(); err == nil {
		data = val
	}
	return data
}
