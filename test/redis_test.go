package test

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	// 连接 redis-server
	// 创建连接
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Printf("redis.Dial() error:%v", err)
		return
	}
	// 关闭连接
	defer c.Close()
	stringSet(c)
}

func stringSet(conn redis.Conn) {
	replySet, err := conn.Do("SET", "key1", "value1")
	if err != nil {
		fmt.Println("SET error: ", err)
	}
	fmt.Println(replySet)
}
