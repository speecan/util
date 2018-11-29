package redisclient

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/speecan/config"
)

var (
	// Pool redis connections
	Pool *redis.Pool
	conf *config.Redis
)

// Init redis pool
func Init(cs *config.Redis) *redis.Pool {
	server := fmt.Sprintf("%s:%d", cs.Host, cs.Port)
	if Pool != nil {
		Pool.Close()
	}
	p := &redis.Pool{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: 0,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if cs.UseAuth {
				if _, err := c.Do("AUTH", cs.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", cs.DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	conf = cs
	Pool = p
	return Pool
}
