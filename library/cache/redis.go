package cache

import (
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

func (c *RedisClient)init() {
	host 		:= os.Getenv("REDIS_HOST")
	password 	:= os.Getenv("REDIS_PASSWORD")
	db 			:= os.Getenv("REDIS_DATABASE")
	port		:= os.Getenv("REDIS_PORT")

	dbInt, _ := strconv.Atoi(db)

	host = host + ":" + port

	redisInstance := redis.NewClient(&redis.Options{
		Addr: host,
		Password: password,
		DB: dbInt,
	})

	pong, err := redisInstance.Ping().Result()
	if err != nil || pong != "PONG" {
		panic("redis server connect exception, " + err.Error())
	}

	c.Client = redisInstance
}

func (c *RedisClient) Set(key string, data string, ttl int) bool {
	var duration time.Duration

	if ttl == 0 {
		duration = time.Duration(0) * time.Second
	} else {
		duration = time.Duration(ttl) * time.Second
	}

	c.Client.Set(key, data, duration)

	return true
}

func (c *RedisClient) Get(key string) string {
	data, _ := c.Client.Get(key).Result()
	return data
}

func (c *RedisClient) SAdd(key string, member []byte) {
	c.Client.SAdd(key, member)
}

func (c *RedisClient) SCard(key string) int {
	count, _ := c.Client.SCard(key).Result()
	return int(count)
}

func (c *RedisClient) SMembers(key string) []string {
	list, _  := c.Client.SMembers(key).Result()
	return list
}