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

	var val string
	val, _ = rclient.Get("mykey_1")
	io.WriteString(w, "mykey_1 = "+val+"\n")

	val, _ = rclient.Get("mykey_2")
	io.WriteString(w, "mykey_1 = "+val+"\n")

	val, _ = rclient.Get("mykey_3")
	io.WriteString(w, "mykey_1 = "+val+"\n")
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
