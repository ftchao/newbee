package controllers

import (
	"io"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func RedisGet(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Redis, get!\n")
}
