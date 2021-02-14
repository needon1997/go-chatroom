package util

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

const addr = "192.168.111.128:6379"

func NewPool() (pool *redis.Pool) {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
