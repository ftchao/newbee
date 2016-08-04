package main

import (
	"fmt"
	"github.com/ftchao/meego"
	"github.com/pelletier/go-toml"
	"newbee/controllers"
	"reflect"
)

func main() {
	fmt.Print("BetaCat Start Port: 9001\n")
	// meego.Run(":9001")
}

func init() {

	init_conf()

	meego.Router("/hello", controllers.HelloServer)
	meego.Router("/redis/get", controllers.RedisGet)
	meego.Router("/redis/set", controllers.RedisSet)
}

// 读取配置文件测试
func init_conf() {

	config, err := toml.LoadFile("conf/main.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	var key string

	key = "redis.auth.user"
	if config.Has(key) {
		fmt.Printf("%s = %s\n", key, config.Get(key))
	}

	key = "redis.connect.timeout"
	if config.Has(key) {
		fmt.Printf("%s = %d\n", key, config.Get(key))
	}

	key = "redis.connect.timeout"
	if config.Has(key) {
		to := config.Get(key)

		// fmt.Printf("\n%#v\n", to)
		fmt.Printf("%s, %s = %d\n", reflect.TypeOf(to), key, to)
	}

	/**
	 * 分割线
	 */
	// redis_auth_user := config.Get("redis.auth.user").(string)
	// fmt.Printf("reids.auth.user = %s\n", redis_auth_user)

	// redis_connect_timeout := config.Get("redis.connect.timeout")
	// fmt.Printf("reids.connect.timeout = %d\n", redis_connect_timeout)

	// redis_read_timeout := config.Get("redis.read.timeout")
	// fmt.Printf("reids.read.timeout = %d\n", redis_read_timeout)

	// redis_write_timeout := config.Get("redis.write.timeout")
	// fmt.Printf("reids.write.timeout = %d\n", redis_write_timeout)

	// configTree := config.Get("redis").(*toml.TomlTree)
	// fmt.Println("User position: %v", configTree.GetPosition("connect.timeout"))

	// // use a query to gather elements without walking the tree
	// fmt.Print("====================================\n")

	// redis_server := config.Get("redis.server").([]interface{})
	// if redis_server != nil {
	// 	for _, v := range redis_server {
	// 		fmt.Printf("ar: %s\n", v.(string))
	// 	}
	// }

	// reiis_server = config.Get("redis.server")

	// // retrieve data directly
	// user := config.Get("postgres.user").(string)
	// password := config.Get("postgres.password").(string)

	// // or using an intermediate object
	// configTree := config.Get("postgres").(*toml.TomlTree)
	// user = configTree.Get("user").(string)
	// password = configTree.Get("password").(string)
	// fmt.Println("User is ", user, ". Password is ", password)

	// // show where elements are in the file
	// fmt.Println("User position: %v", configTree.GetPosition("user"))
	// fmt.Println("Password position: %v", configTree.GetPosition("password"))

	// // use a query to gather elements without walking the tree
	// results, _ := config.Query("$..[user,password]")
	// for ii, item := range results.Values() {
	// 	fmt.Println("Query result %d: %v", ii, item)
	// }
}
