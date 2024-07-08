package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"love_knot/utils/encrypt"
	"time"
)

type JwtTokenStorage struct {
	redis *redis.Client
}

func NewTokenSessionStorage(redis *redis.Client) *JwtTokenStorage {
	return &JwtTokenStorage{redis: redis}
}

func (j *JwtTokenStorage) row(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", encrypt.Md5(token))
}

func (j *JwtTokenStorage) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return j.redis.WithContext(ctx).Set(j.row(token), 1, exp).Err()
}

func (j *JwtTokenStorage) IsBlackList(ctx context.Context, token string) bool {
	return j.redis.WithContext(ctx).Get(j.row(token)).Val() != ""
}
