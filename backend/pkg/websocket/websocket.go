package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// define an Upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// 检查连接来源
	// 为了允许从React发起请求
	// 但我们这里选择不作检查
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	return ws, err
}
