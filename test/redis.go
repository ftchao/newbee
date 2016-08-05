package main

import (
	"fmt"
	"gopkg.in/redis.v4"
)

type RedisClient struct {
	rc *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.255.56:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err == nil {
		return nil
	}

	return client
}

func RedisGet() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.255.56:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	val, err := client.Get("mykey").Result()
	fmt.Println(val, err)
}

func (r *RedisClient) Get(key, value string) {

}
