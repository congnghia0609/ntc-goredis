package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Redis configuration
const (
	RAW_URL    = "redis://127.0.0.1:6379"
	MAX_IDLE   = 5
	MAX_ACTIVE = 10
)

// Redis pool
var redisPool *redis.Pool

// Redis struct
type URedis struct {
	// Redis dial url
	rawUrl string

	// Maximum number of idle connection
	maxIdle int

	// Maximun number of active connection
	maxActive int
}

// Use default redis configuration
func (r *URedis) UseDefaultConfig() {
	r.rawUrl = RAW_URL
	r.maxIdle = MAX_IDLE
	r.maxActive = MAX_ACTIVE
}

// Get redis pool
func (r *URedis) getPool() *redis.Pool {
	if redisPool == nil {
		r.UseDefaultConfig()
		redisPool = &redis.Pool{
			MaxIdle:     r.maxIdle,
			MaxActive:   r.maxActive,
			Wait:        true,
			IdleTimeout: 120 * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(r.rawUrl)
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < 10*time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return redisPool
}

// Put key-value into redis.
func (r *URedis) Put(key string, value string) (string, error) {
	conn := r.getPool().Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("SET", key, value))
	return reply, err
}

// Get key-value from Redis.
func (r *URedis) Get(key string) (string, error) {
	conn := r.getPool().Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	return string(reply), err
}

// Delete key-value from Redis
func (r *URedis) Delete(key string) (string, error) {
	conn := r.getPool().Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("DEL", key))
	return string(reply), err
}

func exampleURedis() {
	redis := new(URedis)
	key1 := "rkey1"
	value1 := "rvalue1"
	fmt.Println("PUT: key=", key1, "value=", value1)
	redis.Put(key1, value1)
	reply, _ := redis.Get(key1)
	fmt.Println("GET: key=", key1, "reply=", reply)
	del, _ := redis.Delete(key1)
	fmt.Println("DEL: key=", key1, "reply=", del)
}
