package redis

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Paschalolo/reddit-recipie-aggregator/internal/repository"
	"github.com/Paschalolo/reddit-recipie-aggregator/pkg"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cache *redis.Client // change to private
	db    repository.Repository
}

var _ repository.CacheRepo = (*Redis)(nil)

func NewRedis(db repository.Repository) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
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
		duration := time.Minute * 20
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

func (c *Redis) PutCookie(ctx context.Context, user string, token *pkg.CookieAuthUser) error {
	duration := time.Until(token.Expires)
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	log.Println("putting cookie")
	c.cache.Set(ctx, user, data, duration)
	return nil
}
func (c *Redis) GetCookie(ctx context.Context, user string) (*pkg.CookieAuthUser, error) {
	var cookie pkg.CookieAuthUser
	value, err := c.cache.Get(ctx, user).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(value), &cookie); err != nil {
		return nil, err
	}
	return &cookie, nil
}
func (c *Redis) DeleteCookie(ctx context.Context, user string) {
	c.cache.Del(ctx, user)
}
