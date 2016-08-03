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
	meego.Router("/hello", controllers.HelloServer)
	meego.Router("/redis/get", controllers.RedisGet)
	meego.Router("/redis/set", controllers.RedisSet)
}
