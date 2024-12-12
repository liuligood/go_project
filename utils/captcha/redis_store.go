package captcha

import (
	"context"
	"crmeb_go/define"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
	"time"
)

type RedisStore struct {
	redisClient redis.UniversalClient
	mtx         sync.Mutex
	ctx         context.Context
}

func (r *RedisStore) Set(id string, value string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	key := define.RedisCaptchaKey + id

	result := r.redisClient.SetNX(r.ctx, key, strings.ToLower(value), time.Second*60)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (r *RedisStore) Get(id string, clear bool) string {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	key := define.RedisCaptchaKey + id
	result := r.redisClient.Get(r.ctx, key)
	if result.Err() != nil {
		return ""
	}

	if clear {
		go r.redisClient.Del(r.ctx, key)
	}

	return result.Val()
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	key := define.RedisCaptchaKey + id
	result := r.redisClient.Get(r.ctx, key)
	if result.Err() != nil {
		return false
	}

	storeAnswer := result.Val()
	if clear {
		go r.redisClient.Del(r.ctx, key)
	}

	return answer == storeAnswer
}

func NewRedisStore(redisClient redis.UniversalClient, ctx context.Context) *RedisStore {
	return &RedisStore{
		redisClient: redisClient,
		mtx:         sync.Mutex{},
		ctx:         ctx,
	}
}
