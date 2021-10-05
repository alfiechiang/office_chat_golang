package main

import (
	// "fmt"
	// "log"
	// "net/http"
	"github.com/polaris1119/chatroom/global"
	"github.com/polaris1119/chatroom/server"
)



func init() {
	global.Init()
}

func main() {
	server.RegisterHandle()
}
