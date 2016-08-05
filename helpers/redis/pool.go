package redis

import (
	"fmt"
)

type RPool struct {
	s int
}

var rpool *RPool

func getConnect() (s int) {
	if rpool == nil {
		rpool = &RPool{0}
	} else {
		rpool.s += 1
	}

	defer func() {
		fmt.Print(">>>>>>>>>>>>>>>> defer \n")
	}()

	return rpool.s
}

// func (r *RPool) newConnect() (s int) {
// 	r.s += 1
// 	return r.s
// }
