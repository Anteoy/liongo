package iredigo

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

const (
	redisServerAdress = "localhost:6379"
	redisPasswd       = "123"
)

var (
	RedisPool *redis.Pool
)

func init() {
	log.Println("redisPoll begin init")
	RedisPool = &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10000,
		IdleTimeout: time.Duration(5) * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServerAdress)
			if err != nil {
				log.Printf("Dail redis server %s %v", redisServerAdress, err)
				return nil, err
			}
			if _, err := c.Do("AUTH", redisPasswd); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("PING"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}
