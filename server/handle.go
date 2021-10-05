package server

import (
	"log"
	"net/http"

	"github.com/polaris1119/chatroom/logic"
	"github.com/rs/cors"
)

var (
	addr = ":2021"

// 	banner = `
//     ____              _____
//    |    |    |   /\     |
//    |    |____|  /  \    |
//    |    |    | /----\   |
//    |____|    |/      \  |

// Go语言编程之旅 —— 一起用Go做项目：ChatRoom，start on：%s

// `
)

func RegisterHandle() {
	// 广播消息处理
	go logic.Broadcaster.Start()

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandleFunc)
	mux.HandleFunc("/user_list", userListHandleFunc)
	mux.HandleFunc("/ws", WebSocketHandleFunc)
	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(addr, handler))

}
