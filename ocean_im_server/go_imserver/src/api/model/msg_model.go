package model

import (
	"fmt"
	"oceanim/go_server/src/api/enum"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/common/db"
	"strconv"
	"time"
)

type ChatMsg struct {
	MsgId       string    `db:"msg_id" json:"msg_id"`
	SendId      string    `db:"send_id" json:"send_id"`
	ReceiveId   string    `db:"receive_id" json:"receive_id"`
	SendTime    time.Time `db:"send_time" json:"send_time"`
	SignType    int       `db:"sign_type" json:"sign_type"`
	MsgType     int       `db:"msg_type" json:"msg_type"`
	ContentType int       `db:"content_type" json:"content_type"`
	Content     string    `db:"content" json:"content"`
	DelFlag     int       `db:"del_flag" json:"del_flag"`
}

func SaveMsg(data *MsgData) (*ChatMsg, error) {
	var chatMsg ChatMsg
	mid := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(time.Now().Unix())
	mid = "msg_" + mid + data.SendId + data.ReceiveId
	chatMsg.MsgId = mid
	chatMsg.SendId = data.SendId
	chatMsg.ReceiveId = data.ReceiveId
	chatMsg.ContentType = data.ContentType
	chatMsg.Content = data.Content
	chatMsg.SendTime = time.Now()
	chatMsg.SignType = enum.MsgNotRead.SignType
	chatMsg.MsgType = enum.MsgTypeText.Type
	chatMsg.DelFlag = enum.MsgNotDel.Flag
	db.InitMysqlDB()
	_, e := db.Db.Exec("INSERT INTO chat_msg(msg_id,send_id,receive_id,msg_type,sign_type,content_type,content,send_time,del_flag) values(?,?,?,?,?,?,?,?,?)",
		&chatMsg.MsgId,
		&chatMsg.SendId,
		&chatMsg.ReceiveId,
		&chatMsg.MsgType,
		&chatMsg.SignType,
		&chatMsg.ContentType,
		&chatMsg.Content,
		&chatMsg.SendTime,
		&chatMsg.DelFlag)

	if e != nil {
		fmt.Println("exec failed, ", e)
		return nil, errorcode.DataErr
	}
	defer db.Db.Close()
	return &chatMsg, nil
}

func SelectMsg(uid string) ([]ChatMsg, error) {
	db.InitMysqlDB()
	var chatMessages []ChatMsg

	rows, e := db.Db.Query("select msg_id,send_id,receive_id,msg_type,sign_type,content_type,content,send_time from chat_msg where  send_id=? or receive_id=? and del_flag=0 order by send_time desc ", uid, uid)
	if e != nil {
		if fmt.Sprint(e) == "sql: no rows in result set" {
			return nil, errorcode.FriendReqSendsNo
		}
		return nil, errorcode.DataErr
	}

	defer db.Db.Close()

	for rows.Next() {
		var chatMsg ChatMsg
		err := rows.Scan(
			&chatMsg.MsgId,
			&chatMsg.SendId,
			&chatMsg.ReceiveId,
			&chatMsg.MsgType,
			&chatMsg.SignType,
			&chatMsg.ContentType,
			&chatMsg.Content,
			&chatMsg.SendTime)
		if err != nil {
			fmt.Println(err)
			return nil, errorcode.DataErr
		}
		chatMessages = append(chatMessages, chatMsg)
	}

	return chatMessages, nil
}

func UpdateMsgSign(msgId string) error {
	db.InitMysqlDB()
	_, err := db.Db.Exec("update chat_msg set sign_type = 1 where msg_id=?", msgId)
	if err != nil {
		fmt.Println("sql exec error", err)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return nil
}

func UpdateMsgDelFlag(id1 string, id2 string) error {
	db.InitMysqlDB()
	_, err1 := db.Db.Exec("update chat_msg set del_flag = 1 where send_id=? and receive_id=?", id1, id2)
	_, err2 := db.Db.Exec("update chat_msg set del_flag = 1 where send_id=? and receive_id=?", id2, id1)
	if err1 != nil || err2 != nil {
		fmt.Println("sql exec error", err1, err2)
		return errorcode.DataErr
	}
	defer db.Db.Close()
	return nil
}
