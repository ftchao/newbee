package main

import (
	"github.com/ftchao/meego"
	"newbee/controllers"
)

func main() {
	meego.Run()
}

func init() {
	meego.Router("/hello", controllers.HelloServer)
	meego.Router("/redis/get", controllers.RedisGet)
}
