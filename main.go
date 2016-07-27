package main

import (
	"github.com/ftchao/meego"
	"newbee/controllers"
)

func main() {
	meego.Run(":9001")
}

func init() {
	meego.Router("/hello", controllers.HelloServer)
	meego.Router("/redis/get", controllers.RedisGet)
}
