package socket_io

import (
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/model"
	"time"
)

func NewSocketClientHub() model.SocketClientHub {
	return model.SocketClientHub{SocketClients: make(map[string]*model.SocketClient)}
}

func NewSocketServerIO(options *engineio.Options) *model.SocketServerIO {
	socketServerIO := socketio.NewServer(options)
	return &model.SocketServerIO{Server: socketServerIO}
}

func InitSocketIoWs(hub *model.SocketClientHub) *model.SocketServerIO {
	server := NewSocketServerIO(nil)
	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		return nil
	})

	server.OnEvent("/", "register", func(s socketio.Conn, msg map[string]interface{}) {
		var socketClient model.SocketClient
		socketClient.Uid = msg["uid"].(string)
		socketClient.Platform = msg["platform"].(string)
		socketClient.LastActTime = time.Now()
		socketClient.Socket = s
		hub.SocketClients[socketClient.Uid] = &socketClient
		s.Join(socketClient.Uid)
		return
	})

	server.OnEvent("/", "send", func(s socketio.Conn, msg map[string]interface{}) {
		//fmt.Println("接收到消息:", msg)
		//fmt.Println("发送者的房间号")
		//fmt.Println(s.Rooms())
		tmp, _ := json.Marshal(msg["data"])
		//fmt.Println(string(tmp))
		//var socketMsg model.SocketMsg
		var msgData model.MsgData
		reply := make(map[string]string)
		//t, _ := msg["type"].(int)
		_ = json.Unmarshal(tmp, &msgData)
		if msgData.Content == "" {
			return
		}
		//socketMsg.Type = t
		//socketMsg.Data = &msgData

		fromRoomId := msgData.SendId
		toRoomId := msgData.ReceiveId
		var f model.Friend
		f.OwnerId = fromRoomId
		f.FriendId = toRoomId
		isFriend := model.CheckIsFriend(&f)
		if isFriend != "true" {
			reply["code"] = "30002"
			reply["send_res"] = errorcode.MsgSendNotFriend.Message
			server.BroadcastToRoom("/", fromRoomId, "sendRes", reply)
			return
		}
		//fmt.Println(toRoomId)
		socketClient := hub.SocketClients[msgData.ReceiveId]
		//fmt.Println("接受者的房间号")
		//fmt.Println(socketClient.Socket.Rooms())

		sMsg, err := model.SaveMsg(&msgData)
		if msgData.SendId == msgData.ReceiveId {
			return
		}
		if sMsg != nil {
			if err != nil {
				code, errMsg := errorcode.ParserErr(err)
				reply["code"] = string(code)
				reply["send_res"] = errMsg
				server.BroadcastToRoom("/", fromRoomId, "sendRes", reply)
			}
			if socketClient != nil {
				//fmt.Println(socketClient)
				server.BroadcastToRoom("/", toRoomId, "singChat", &sMsg)
				reply["code"] = "00000"
				reply["send_res"] = errorcode.Ok.Message
				server.BroadcastToRoom("/", fromRoomId, "sendRes", reply)
			} else {
				reply["code"] = "30003"
				reply["send_res"] = errorcode.MsgSendNotOnline.Message
				server.BroadcastToRoom("/", fromRoomId, "sendRes", reply)
			}
		}
		return
	})

	server.OnEvent("/", "logout", func(s socketio.Conn, uid string) {
		if hub.SocketClients[uid] != nil {
			delete(hub.SocketClients, uid)
			s.LeaveAll()
			s.Close()
			return
		}
		return

	})

	server.OnEvent("/", "ping", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		_ = s.Close()
		return last
	})

	server.OnEvent("/", "heartBeat", func(s socketio.Conn, tmp map[string]interface{}) {
		uid := tmp["uid"].(string)
		var timeLayoutStr = "2006-01-02 15:04:05"
		lastActTime, _ := time.ParseInLocation(timeLayoutStr, tmp["last_act_time"].(string), time.Local)
		reply := make(map[string]string)
		s.Join(uid)
		//fmt.Println("lastActTime")
		//fmt.Println(lastActTime)
		//fmt.Println("hub.SocketClients[uid].LastActTime")
		//fmt.Println(hub.SocketClients[uid].LastActTime)
		//fmt.Println(lastActTime.Sub(hub.SocketClients[uid].LastActTime).Seconds())
		var socketClient model.SocketClient
		if hub.SocketClients[uid].Socket == s &&
			lastActTime.Sub(hub.SocketClients[uid].LastActTime).Seconds() < 10 {
			socketClient.Socket = s
			socketClient.LastActTime = lastActTime
			socketClient.Uid = uid
			socketClient.Platform = hub.SocketClients[uid].Platform
			hub.SocketClients[uid] = &socketClient
			reply["status"] = "heart success"
			server.BroadcastToRoom("/", uid, "heartRes", reply)
			return
		} else {
			reply["status"] = "heart fail"
			server.BroadcastToRoom("/", uid, "heartRes", reply)
			return
		}
	})

	server.OnEvent("/", "readMsg", func(s socketio.Conn, msgId string) {
		err := model.UpdateMsgSign(msgId)
		if err != nil {
			code, msg := errorcode.ParserErr(err)
			fmt.Println(code, msg)
		}
		return
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
		for _, client := range hub.SocketClients {
			if client.Socket == s {
				delete(hub.SocketClients, client.Uid)
				s.LeaveAll()
				s.Close()
				return
			}
		}
		return
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed:", reason)
		for _, client := range hub.SocketClients {
			if client.Socket == s {
				delete(hub.SocketClients, client.Uid)
				s.LeaveAll()
				s.Close()
				return
			}
		}
		return
	})

	return server
}

var SocketHub = NewSocketClientHub()
var Server = InitSocketIoWs(&SocketHub)
