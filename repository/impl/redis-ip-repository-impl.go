package implements

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

type RedisIpRepositoryImpl struct {
	client *redis.Client
}

var (
	redisIpRepoInstance *RedisIpRepositoryImpl
	redisIpRepoOnce     sync.Once
)

func GetRedisIpRepository(addr, password string, db int) *RedisIpRepositoryImpl {
	redisIpRepoOnce.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})
		redisIpRepoInstance = &RedisIpRepositoryImpl{client: rdb}
	})
	return redisIpRepoInstance
}

func (r *RedisIpRepositoryImpl) FindByHost(host string) ([]string, error) {
	ctx := context.Background()
	ips, err := r.client.LRange(ctx, host, 0, -1).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("host not found: %s", host)
	} else if err != nil {
		return nil, err
	}
	return ips, nil
}

func (r *RedisIpRepositoryImpl) Create(host string, newIps []string) error {
	ctx := context.Background()
	err := r.client.Del(ctx, host).Err()
	if err != nil {
		return err
	}
	err = r.client.RPush(ctx, host, newIps).Err()
	return err
}

func (r *RedisIpRepositoryImpl) Update(host string, updateIps []string) error {
	ctx := context.Background()
	exists, err := r.client.Exists(ctx, host).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return fmt.Errorf("host not found: %s", host)
	}
	err = r.client.Del(ctx, host).Err()
	if err != nil {
		return err
	}
	err = r.client.RPush(ctx, host, updateIps).Err()
	return err
}

func (r *RedisIpRepositoryImpl) Delete(host string) error {
	ctx := context.Background()
	err := r.client.Del(ctx, host).Err()
	return err
}
