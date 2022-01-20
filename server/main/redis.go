package main

import (
	"time"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initPool(adress string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", adress)
		},
	}
}
