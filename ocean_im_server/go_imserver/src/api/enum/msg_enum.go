/****
**** @Author: OCEAN GZY
**** @Date:   2022/2/27 21:21
****/

package enum

type MsgType struct {
	Type    int    `json:"type"`
	Comment string `json:"comment"`
}

type MsgSignType struct {
	SignType int    `json:"sign_type"`
	Comment  string `json:"comment"`
}

type MsgDelFlag struct {
	Flag    int    `json:"flag"`
	Comment string `json:"comment"`
}

var (
	MsgTypePing    = &MsgType{Type: 0, Comment: "ping"}
	MsgTypeOnline  = &MsgType{Type: 1, Comment: "上线"}
	MsgTypeOffline = &MsgType{Type: 2, Comment: "下线"}
	MsgTypeText    = &MsgType{Type: 3, Comment: "文本消息"}
	MsgTypeImage   = &MsgType{Type: 4, Comment: "图片消息"}

	MsgIsRead  = &MsgSignType{SignType: 1, Comment: "已读"}
	MsgNotRead = &MsgSignType{SignType: 0, Comment: "未读"}

	MsgIsDel  = &MsgDelFlag{Flag: 1, Comment: "已删除"}
	MsgNotDel = &MsgDelFlag{Flag: 0, Comment: "未删除"}
)
