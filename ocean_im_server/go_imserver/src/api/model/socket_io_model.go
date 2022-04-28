package model

import (
	socketio "github.com/googollee/go-socket.io"
	"time"
)

type SocketServerIO struct {
	*socketio.Server
}

type SocketMsg struct {
	Type int      `json:"type"`
	Data *MsgData `json:"data"`
}

type MsgData struct {
	Platform    string `json:"platform"`
	SendId      string `json:"send_id"`
	GroupId     string `json:"group_id"`
	ReceiveId   string `json:"receive_id"`
	ContentType int    `db:"content_type"`
	Content     string `json:"content"`
}
type SocketClient struct {
	Platform string        `json:"platform"`
	Uid      string        `json:"uid"`
	Socket   socketio.Conn `json:"socket"`
	LastActTime time.Time `json:"last_act_time"`
}

type SocketClientHub struct {
	SocketClients map[string]*SocketClient
}
