package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cache *redis.Client // change to private
	db    repository.Repository
}

var _ repository.CacheRepo = (*Redis)(nil)

func NewRedis(db repository.Repository) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	status := rdb.Ping(context.Background())
	log.Println(status)
	log.Println("Redis cache added ")
	return &Redis{cache: rdb, db: db}
}

// Get returns a string of cached data and error
// if Cache is not available Get will automatically make a call to the database
func (c *Redis) Get(ctx context.Context, key string) (string, error) {
	r, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("requet to mongo db ")
		result, err := c.db.GetRecipe(ctx)
		if err != nil {
			return "", err
		}
		data, err := json.Marshal(result)
		if err != nil {
			return "", err
		}
		duration := time.Minute * 10
		log.Println("Putting data in the cache")
		c.cache.Set(ctx, key, data, duration)
		return "", repository.ErrNotInCache
	}
	if err != nil {
		return "", repository.ErrCache
	}
	return r, nil
}

func (c *Redis) Delete(ctx context.Context, key string) {
	log.Println("Deleteing cache")
	c.cache.Del(ctx, key)
}
