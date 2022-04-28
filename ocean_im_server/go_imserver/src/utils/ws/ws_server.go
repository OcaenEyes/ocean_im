/****
**** @Author: OCEAN GZY
**** @Date:   2022/2/27 20:29
****/

package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func MyWsServer(hub *MyWsHub, c *gin.Context) {
	initWs(hub, c)
}


var upgrade = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func initWs(hub *MyWsHub, c *gin.Context) {
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	client := &MyWsClient{Hub: hub, Socket: ws, Send: make(chan []byte, 256)}
	client.Hub.ChRegister <- client

	go client.readMsg()
	go client.writeMsg()
}