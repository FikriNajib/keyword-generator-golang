package infrastructure

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.elastic.co/apm/module/apmgoredisv8"
	"keyword-generator/src/config"
	"log"
	"time"
)

type RedisDatabase struct {
	Client *redis.Client
}

var AppRedis *RedisDatabase

func init() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	var redisDb RedisDatabase
	address := config.Config.GetString("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.Config.GetString("REDIS_PASSWORD"), // no password set
		DB:       config.Config.GetInt("REDIS_DB"),          // use default DB
	})
	client.AddHook(apmgoredisv8.NewHook())
	redisDb.Client = client
	AppRedis = &redisDb
}

func (r *RedisDatabase) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisDatabase) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {

	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.Client.Set(ctx, key, val, ttl).Err()
}

func (r *RedisDatabase) Close() error {
	return r.Client.Close()
}

func (r *RedisDatabase) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}
