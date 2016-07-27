package main

import (
	"fmt"
	"github.com/ftchao/meego"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"newbee/controllers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, foo!\n")
}

func json(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "{\"role_name\":\"开发\",\"p_size\":5000,\"system_key\":\"pdms\"}\n")
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handler)
	r.HandleFunc("/foo", foo)
	r.HandleFunc("/json", json)

	// callback fun
	// r.HandleFunc("/order/create", &controllers.OrderController{}, "*:Get")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":9001", r))
}

func init() {
	// redis 测试
	// RedisGet()

	var rc controllers.RedisAction
	rc.Get()

	// rc := RedisClient()
	// rc.Get()
	// rc.Set()

	fmt.Println("man init")
}
