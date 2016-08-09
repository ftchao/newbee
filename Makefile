# GOPATH:=$(CURDIR)
# export GOPATH:=$(CURDIR)

# TARGET = bin/newbee

# all: build

fmt:
	# gofmt -l -w -s

dep:fmt
	go get gopkg.in/redis.v4
	go get github.com/ftchao/meego
	go get github.com/pelletier/go-toml

run: dep
	go run main.go controllers helpers
