package redis

import (
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

const (
	failPrefix = `failCount_`
)

func SetFailCount(payloadKey string, count int, timeout int) (err error) {
	con := pool.Get()
	defer con.Close()
	countStr := strconv.Itoa(count)
	key := failPrefix + payloadKey
	_, err = con.Do("setex", key, timeout, countStr)
	if err != nil {
		fmt.Println("redis SetFailCount err,", err)
		return
	}
	return
}

func GetFailCount(payloadKey string) (count int, err error) {
	con := pool.Get()
	defer con.Close()
	key := failPrefix + payloadKey
	count, err = redis.Int(con.Do("get", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			count = 0
			return
		}
		fmt.Println("redis GetFailCount err,", err)
		return
	}
	return
}
