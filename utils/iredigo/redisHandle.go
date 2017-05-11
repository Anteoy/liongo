package iredigo

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

func Do(commandName string, args ...interface{}) (interface{}, error) {
	c := RedisPool.Get()
	defer c.Close()

	result, err := c.Do(commandName, args...)
	if err != nil {
		log.Panic(err)
	}
	return result, err
}

func Incr(name string) int {
	c := RedisPool.Get()
	defer c.Close()

	result, err := redis.Int(c.Do("INCR", name))
	if err != nil {
		log.Panic(err)
	}
	return result
}

func Get(name string) string {
	c := RedisPool.Get()
	defer c.Close()

	return GetFromConn(c, name)
}

func GetBytes(name string) []byte {
	c := RedisPool.Get()
	defer c.Close()

	return GetFromConnBytes(c, name)
}

func GetFromConn(c redis.Conn, name string) string {

	result, err := redis.String(c.Do("GET", name))
	if err != nil {
		log.Panicln("redis GET " + name + " is error")
	}

	return result
}

func GetFromConnBytes(c redis.Conn, name string) []byte {

	result, err := redis.Bytes(c.Do("GET", name))
	if err != nil {
		log.Panicln("redis GET " + name + " is error")
	}

	return result
}

func Set(name string, value string, timeout int64) int {
	c := RedisPool.Get()
	defer c.Close()

	return SetFromConn(c, name, value, timeout)
}
func SetBytes(name string, value []byte, timeout int64) int {
	c := RedisPool.Get()
	defer c.Close()

	return SetFromConnBytes(c, name, value, timeout)
}

func SetFromConn(c redis.Conn, name string, value string, timeout int64) int {

	_, err := c.Do("SET", name, value)
	_, err = c.Do("EXPIRE", name, timeout)
	if err != nil {
		return 0
	}
	return 1
}
func SetFromConnBytes(c redis.Conn, name string, value []byte, timeout int64) int {

	_, err := c.Do("SET", name, value)
	_, err = c.Do("EXPIRE", name, timeout)
	if err != nil {
		return 0
	}
	return 1
}