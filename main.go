package main

import (
	"fmt"
	"github.com/ftchao/meego"
	"newbee/controllers"
	"newbee/helpers/toml"
)

func main() {
	fmt.Print("BetaCat Start Port: 9001\n")
	meego.Run(":9001")
}

func init() {

	// init_conf()

	meego.Router("/hello", controllers.HelloServer)
	meego.Router("/redis/get", controllers.RedisGet)
	meego.Router("/redis/set", controllers.RedisSet)
}

// 读取配置文件测试
func init_conf() {

	conf, err := toml.LoadFile("conf/main.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	var key string

	key = "redis.auth.user"
	fmt.Printf("%s = %s\n", key, conf.Get(key))

	key = "redis.connect.timeout"
	fmt.Printf("%s = %d\n", key, conf.Get(key))

	redis_server := conf.Get("redis.server").([]interface{})
	if redis_server != nil {
		for _, v := range redis_server {
			fmt.Printf("ar: %s\n", v.(string))
		}
	}

}
