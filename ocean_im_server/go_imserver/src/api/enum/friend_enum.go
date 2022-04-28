/****
**** @Author: OCEAN GZY
**** @Date:   2022/2/20 19:52
****/

package enum

type FriendReqFlag struct {
	Flag    int  `json:"flag"`
	Comment string `json:"comment"`
}

type FriendFlag struct {
	Flag    int  `json:"flag"`
	Comment string `json:"comment"`
}

var (
	Todo   = &FriendReqFlag{Flag: 0, Comment: "未处理"}
	Accept = &FriendReqFlag{Flag: 1, Comment: "已接受"}
	Refuse = &FriendReqFlag{Flag: 2, Comment: "已拒绝"}
	Black  = &FriendReqFlag{Flag: 3, Comment: "已拉黑"}
	Ignore  = &FriendReqFlag{Flag: 4, Comment: "已忽略"}
	NoFriend = &FriendFlag{Flag: 0, Comment: "是好友"}
	IsFriend = &FriendFlag{Flag: 1, Comment: "不是好友"}
)
