package redis

import (
	// "fmt"
	r "github.com/garyburd/redigo/redis"
	// "newbee/helper/toml"
	// "sync"
	// "time"
)

// var rpool_ctx sync.One

type RClient struct {
	pool *r.Pool
}

func Connect(server string) *RClient {

	// 	// connect to redis server
	// 	redis_server := config.get()
	// 	for _, server := range redis_server {
	// 		return &RClient {
	// 			pool: getConnect(server)
	// 		}
	// 	}

	return &RClient{}
}

// func getConnect(server string) *r.Pool {

// 	return r.Pool{
// 		MaxIdle:     3,
// 		IdleTimeout: 240 * time.Second,
// 		Dial: func(r.Conn, error) {

// 			for i := 0; i < count; i++ {
// 				c, err := r.Dial("tcp", server)
// 				if err != nil {
// 					return nil, err
// 				}
// 				return c, err
// 			}

// 			return nil, errors.New("redispool: cannot connect to any redis server")
// 		},
// 	}
// }

func (r *RClient) Get() string {
	return "sdfsf"
}

// type RClient struct {
// 	client       *r.Client
// 	Redisaddr    *[]string
// 	conntimeout  int
// 	DialTimeout  int
// 	ReadTimeout  int
// 	WriteTimeout int
// }

// const (
// 	connectTry = 10 // 连接重试次数
// )

// func init() {
// 	// 读取配置文件
// 	fmt.Print("redis  client init\n")

// 	// PoolClient()
// }

// func NewClient() (rc *RClient) {
// 	client := r.NewClient(&r.Options{
// 		Addr:     "192.168.255.56:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	pong, err := client.Ping().Result()
// 	if err != nil {
// 		fmt.Println(pong, err) // todo write log
// 		return nil
// 	}

// 	return &RClient{client: client}
// }

// func (r *RClient) Close() {
// 	r.Close()
// }

// func (r *RClient) Get(key string) (string, error) {
// 	val, err := r.client.Get(key).Result()
// 	if err != nil {
// 		return "", err
// 	}

// 	return val, nil
// }

// func (r *RClient) Set(key string, value interface{}, expiration time.Duration) *r.StatusCmd {

// 	if r == nil {
// 		return nil
// 	}

// 	err := r.client.Set(key, value, expiration)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *RClient) SetNX(key string, value interface{}, expiration time.Duration) *r.BoolCmd {
// 	err := r.client.SetNX(key, value, expiration)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
