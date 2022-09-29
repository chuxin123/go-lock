package lock

import (
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

const (
	script = `local key = KEYS[1]
local res = redis.call("GET", key)
if res then
	return redis.call("DEL", key)
else
	return 0
end
`
)

type redisLock struct {
}

var (
	redisLockKey = "redis-lock-key"
	redisClient  *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func NewRedisLock() *redisLock {
	return &redisLock{}
}

func (r *redisLock) lock() bool {

	uniqueId := rand.Int() // 设定owner
	resp := redisClient.SetNX(redisLockKey, uniqueId, 10*time.Second)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		return false
	}
	return true
}

// 正式使用需要验证owner
func (r *redisLock) unlock() bool {
	resp := redisClient.Eval(script, []string{redisLockKey})
	unlockSuccess, err := resp.Result()
	v, ok := unlockSuccess.(bool)
	if err != nil || !ok || !v {
		return false
	}
	return true
}
