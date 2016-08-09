package controllers

import (
	"fmt"
	"github.com/ftchao/meego"
	"io"
	"net/http"
	// "newbee/helpers/redis"
	// "time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func RedisGet(w http.ResponseWriter, req *http.Request) {
	name := meego.ConfGet("test.name").(string)
	io.WriteString(w, "conf mysql.dsn = "+name+"\n")

	size := meego.ConfGet("test.size")
	io.WriteString(w, "conf mysql.dsn = "+fmt.Sprintf("%d", size)+"\n")

	// var key string

	// rc := redis.Connect()

	// key = "mykey_1"
	// v, _ := rc.Get(key)
	// io.WriteString(w, key+": "+v)
}

func RedisSet(w http.ResponseWriter, req *http.Request) {

}
