package controllers

import (
	"fmt"
	"github.com/ftchao/meego"
	// "github.com/ftchao/meego/redis"
	"io"
	"net/http"
	// "newbee/helpers/redis"
	// "time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

/**
 * Redis test
 */
func RedisGet(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func RedisSet(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

/**
 * Cont test
 */
func ConfGet(w http.ResponseWriter, req *http.Request) {
	name := meego.ConfGet("test.name").(string)
	io.WriteString(w, "conf mysql.dsn = "+name+"\n")

	size := meego.ConfGet("test.size")
	io.WriteString(w, "conf mysql.dsn = "+fmt.Sprintf("%d", size)+"\n")

	fruits := meego.ConfGet("test.fruits").([]interface{})
	for _, name := range fruits {
		io.WriteString(w, "fruits: "+name.(string)+"\n")
	}
}
