/**
 *
 * @author nghiatc
 * @since Feb 8, 2018
 */

package nredis

import (
	"fmt"
	"github.com/congnghia0609/ntc-gconf/nconf"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Redis configuration
const (
	RAW_URL    = "redis://127.0.0.1:6379"
	MAX_IDLE   = 5
	MAX_ACTIVE = 10
	TIMEOUT = 120
)

// Redis pool
var redisPool *redis.Pool

func InitPoolConf(name string) *redis.Pool {
	c := nconf.GetConfig()
	maxIdle := c.GetInt(name+".redis.max_idle")
	if maxIdle == 0 {
		maxIdle = MAX_IDLE
	}
	//fmt.Println(maxIdle)
	maxActive := c.GetInt(name+".redis.max_active")
	if maxActive == 0 {
		maxActive = MAX_ACTIVE
	}
	//fmt.Println(maxActive)
	timeout := c.GetInt64(name+".redis.timeout")
	if timeout == 0 {
		timeout = TIMEOUT
	}
	//fmt.Println(timeout)
	idleTimeout := time.Duration(timeout) * time.Second
	rawUrl := c.GetString(name+".redis.url")
	if rawUrl == "" {
		rawUrl= RAW_URL
	}
	//fmt.Println(rawUrl)
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			Wait:        true,
			IdleTimeout: idleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(rawUrl)
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < 5*time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return redisPool
}

func InitPool(rawUrl string, maxIdle int, maxActive int, idleTimeout time.Duration) *redis.Pool {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			Wait:        true,
			IdleTimeout: idleTimeout,
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(rawUrl)
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < 5*time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return redisPool
}

// Get redis pool
func GetPool() *redis.Pool {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:     MAX_IDLE,
			MaxActive:   MAX_ACTIVE,
			Wait:        true,
			IdleTimeout: TIMEOUT * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(RAW_URL)
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < 5*time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return redisPool
}

func GetConnection() redis.Conn {
	return GetPool().Get()
}

// Put key-value into redis.
func Put(key string, value string) (string, error) {
	conn := GetPool().Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("SET", key, value))
	return reply, err
}

// Get key-value from Redis.
func Get(key string) (string, error) {
	conn := GetPool().Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	return string(reply), err
}

// Delete key-value from Redis
func Delete(key string) (string, error) {
	conn := GetPool().Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("DEL", key))
	return string(reply), err
}

func ExampleNRedis() {
	key1 := "rkey1"
	value1 := "rvalue1"
	fmt.Println("PUT: key=", key1, "value=", value1)
	Put(key1, value1)
	reply, _ := Get(key1)
	fmt.Println("GET: key=", key1, "reply=", reply)
	del, _ := Delete(key1)
	fmt.Println("DEL: key=", key1, "reply=", del)
}
