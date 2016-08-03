package controllers

import (
	"io"
	"net/http"
	"newbee/helpers/redis"
	"time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func RedisGet(w http.ResponseWriter, req *http.Request) {
	rclient := redis.NewClient()
	val := rclient.Abcc()

	io.WriteString(w, "Redis, get!\n")
	io.WriteString(w, val+"\n") // todo '+' 效率低
}

func RedisSet(w http.ResponseWriter, req *http.Request) {
	rclient := redis.NewClient()

	val := "slj2i2oi3r" + time.Now().String()
	err := rclient.Set("mykey_1", val, 0)
	if err != nil {
		io.WriteString(w, "set mykey_1 failed\n")
	}

	io.WriteString(w, "set mykey_1 ok\n")
}
