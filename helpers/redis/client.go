package redis

import (
	"fmt"
	r "gopkg.in/redis.v4"
	// "math/rand"
	"time"
)

type RClient struct {
	client       *r.Client
	Redisaddr    *[]string
	conntimeout  int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

const (
	connectTry = 10 // 连接重试次数
)

func init() {
	// 读取配置文件
	fmt.Print("redis  client init\n")

	// PoolClient()
}

func NewClient() (rc *RClient) {
	client := r.NewClient(&r.Options{
		Addr:     "192.168.255.56:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err) // todo write log
		return nil
	}

	return &RClient{client: client}
}

func (r *RClient) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *RClient) Set(key string, value interface{}, expiration time.Duration) *r.StatusCmd {

	if r == nil {
		return nil
	}

	err := r.client.Set(key, value, expiration)
	if err != nil {
		return err
	}

	return nil
}

func (r *RClient) SetNX(key string, value interface{}, expiration time.Duration) *r.BoolCmd {
	err := r.client.SetNX(key, value, expiration)
	if err != nil {
		return err
	}

	return nil
}
