package redis

import "errors"

var (
	RedisError = errors.New("redis error")
)

const (
	KeyPrefix        = "go_mall:"
	KeyExpiredTokens = "expireTokens"
)

func GetKey(key string) string {
	return KeyPrefix + key
}
