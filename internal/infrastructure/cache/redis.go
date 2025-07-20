package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type Cache[T any] struct {
	cache   *cache.Cache
	context context.Context
	ttl     time.Duration
}

func NewCache[T any](redis *redis.Client, ctx context.Context, ttl time.Duration) *Cache[T] {
	return &Cache[T]{
		cache: cache.New(&cache.Options{
			Redis:      redis,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		}),
		context: ctx,
		ttl:     ttl,
	}
}

// adds or updates a value in the cache
func (c *Cache[T]) Set(key string, value T) error {
	item := &cache.Item{
		Key:   key,
		Value: value,
		TTL:   c.ttl,
	}
	return c.cache.Set(item)
}

// adds or updates an array of values in the cache
func (c *Cache[T]) SetArray(key string, values []T) error {
	item := &cache.Item{
		Key:   key,
		Value: values,
		TTL:   c.ttl,
	}
	return c.cache.Set(item)
}

// retrieves a value from the cache
func (c *Cache[T]) Get(key string) (T, error) {
	var value T
	err := c.cache.Get(c.context, key, &value)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return value, nil
		}
		return value, err
	}
	if err != nil {
		return value, err
	}
	return value, nil
}

// retrieves an array of values from the cache
func (c *Cache[T]) GetArray(key string) ([]T, error) {
	var values []T
	err := c.cache.Get(c.context, key, &values)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return values, nil
		}
		return values, err
	}
	if err != nil {
		return values, err
	}
	return values, nil
}

// removes a specific key from the cache
func (c *Cache[T]) Delete(key string) error {
	return c.cache.Delete(c.context, key)
}

// retrieves or stores a value in cache
func (c *Cache[T]) GetOrSet(key string, fn func() (T, error)) (T, error) {
	cached, err := c.Get(key)
	if err == nil {
		return cached, nil
	}
	value, err := fn()
	if err != nil {
		return value, err
	}
	return value, c.Set(key, value)
}

// retrieves or stores an array of values in cache
func (c *Cache[T]) GetOrSetArray(key string, fn func() ([]T, error)) ([]T, error) {
	cached, err := c.GetArray(key)
	if err == nil {
		return cached, nil
	}
	value, err := fn()
	if err != nil {
		return value, err
	}
	return value, c.SetArray(key, value)
}
