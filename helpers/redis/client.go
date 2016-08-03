package redis

import (
	"fmt"
	r "gopkg.in/redis.v4"
	"math/rand"
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

	PoolClient()
}

func PoolClient() {

	ini_conf := []string{
		"192.168.255.56:6371",
		"192.168.255.56:6372",
		"192.168.255.56:6373",
		"192.168.255.56:6374",
	}

	// 连接选择
	for i := 0; i < connectTry; i++ {

		s := rand.Intn(10)
		fmt.Printf("addr: %d,%d, %s\n", i, s, ini_conf[0])

		// for _, v := range ini_conf {
		// 	fmt.Printf("addr: %s\n", v)
		// }
	}

}

func NewClient() *RClient {

	client := r.NewClient(&r.Options{
		// Addr:     "127.0.0.1:6379",
		Addr:     "192.168.255.56:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
		return nil
	}

	rclient := &RClient{
		client: client,
	}

	return rclient
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
		fmt.Print("================================ client nil\n")
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

// test func
func (r *RClient) Abcc() string {
	return time.Now().String()
}
