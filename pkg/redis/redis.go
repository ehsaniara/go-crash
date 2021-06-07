package redis

import (
	"encoding/json"
	"fmt"
	"github.com/ehsaniara/go-crash/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() {
	RedisConn = &redis.Pool{
		MaxIdle:     config.AppConfig.Redis.MaxIdle,
		MaxActive:   config.AppConfig.Redis.MaxActive,
		IdleTimeout: config.AppConfig.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.AppConfig.Redis.Host)
			if err != nil {
				return nil, err
			}
			if config.AppConfig.Redis.Password != "" {
				if _, err := c.Do("AUTH", config.AppConfig.Redis.Password); err != nil {
					err := c.Close()
					if err != nil {
						return nil, err
					}
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	fmt.Printf(" - Redis Started on Host: %s\n", config.AppConfig.Redis.Host)
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Errorf("Redi set %v\n", err)
		}
	}(conn)

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Errorf("Redi Exists %v\n", err)
		}
	}(conn)

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Errorf("Redi Get %v\n", err)
		}
	}(conn)

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Errorf("Redi Delete %v\n", err)
		}
	}(conn)

	return redis.Bool(conn.Do("DEL", key))
}
