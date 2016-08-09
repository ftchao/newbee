package main

import (
	"fmt"
	"github.com/ftchao/meego"
	"newbee/controllers"
)

func main() {
	fmt.Print("BetaCat Start Port: 9001\n")
	meego.Run(":9001")
}

func init() {

	// 读取配置文件
	meego.LoadConf("conf/main.toml")

	// name := meego.ConfGet("test.name").(string)
	// fmt.Printf("============== %s \n", name)

	// size := meego.ConfGet("test.size")
	// fmt.Printf("============== %d \n", size)

	// 路由设置
	meego.Router("/hello", controllers.HelloServer)
	meego.Router("/conf/get", controllers.ConfGet)
	meego.Router("/redis/get", controllers.RedisGet)
	meego.Router("/redis/set", controllers.RedisSet)
}
