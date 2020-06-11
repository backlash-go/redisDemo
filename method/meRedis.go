package method

import (
	"github.com/garyburd/redigo/redis"
	"redisDemo/conf"
)

func SetHashValues(keyFiled []interface{}) error {
	conn := conf.GetRedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("HMSET", keyFiled...)
	return err
}

func SetKeyTtl(key string, expires int) error {
	conn := conf.GetRedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, expires)
	return err
}

func GetHashValues(keyFiled []interface{}) ([]string, error) {
	conn := conf.GetRedisPool().Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HMGET", keyFiled...))
}

func SetEx(key, value string, expires int) error {
	conn := conf.GetRedisPool().Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", key, expires, value)
	return err
}

func GetStringKey(key string) (string, error) {
	conn := conf.GetRedisPool().Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}
