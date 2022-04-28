/****
**** @Author: OCEAN GZY
**** @Date:   2022/2/27 20:30
****/

package ws

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type MyWsMsg struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type MyWsClient struct {
	Addr     string          `json:"addr"`
	Socket   *websocket.Conn `json:"socket"`
	Platform string          `json:"platform"`
	Uid      string          `json:"uid"`
	GroupId  string          `json:"group_id"`
	ToUid    string          `json:"to_uid"`
	Send     chan []byte     `json:"send"`
	Hub      *MyWsHub
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

func (c *MyWsClient) readMsg() {
	defer func() {
		c.Hub.ChUnregister <- c
		c.Socket.Close()
	}()

	c.Socket.SetReadLimit(int64(maxMessageSize))
	c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	c.Socket.SetPongHandler(func(appData string) error {
		c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Errorf("error:%v", err)
			}
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		c.Hub.Broadcast <- msg
	}
}

func (c *MyWsClient) writeMsg() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Socket.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Socket.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.Socket.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}
