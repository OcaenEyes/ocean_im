package friend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanim/go_server/src/api/enum"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/model"
	"oceanim/go_server/src/utils/auth"
	"oceanim/go_server/src/utils/socket_io"
	"time"
)

func SendFriendReq(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var friendReq model.FriendRequest
	err := c.ShouldBindJSON(&friendReq)
	if err != nil || friendReq.ReqId == "" || friendReq.UserId == "" {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.FriendReqParamErr.Code, "message": errorcode.FriendReqParamErr.Message})
		return
	}
	var f model.Friend
	f.OwnerId = friendReq.UserId
	f.FriendId = friendReq.ReqId
	isFriend := model.CheckIsFriend(&f)
	if isFriend == "true" {
		return
	} else if isFriend == "false" {
		friendReq.Flag = enum.Todo.Flag
		friendReq.CreateTime = time.Now()
		err = model.SaveFriendRequest(&friendReq)
		user, _ := model.UserInfoByUid(uid)
		var tmpf model.FriendInfo
		tmpf.Uid = user.Uid
		tmpf.Remark = ""
		tmpf.NickName = user.NickName
		tmpf.Phone = user.Phone
		tmpf.Email = user.Email
		tmpf.Ex = user.Ex
		tmpf.Gender = user.Gender
		tmpf.Birth = user.Birth

		res := make(map[string]interface{})
		res["friendR"] = &friendReq
		res["tmpInfo"] = &tmpf
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		socket_io.Server.BroadcastToRoom("/", friendReq.UserId, "receive_friend_req", res)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.DataErr.Code, "message": errorcode.DataErr.Message})
		return
	}
	return
}

func DueFriendReq(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var friendReq model.FriendRequest
	err := c.ShouldBindJSON(&friendReq)
	if err != nil || friendReq.ReqId == "" || friendReq.UserId == "" {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.FriendReqParamErr.Code, "message": errorcode.FriendReqParamErr.Message})
		return
	}
	switch friendReq.Flag {
	case enum.Accept.Flag:
		var friend model.Friend
		friend.OwnerId = friendReq.UserId
		friend.FriendId = friendReq.ReqId
		friend.FriendFlag = enum.IsFriend.Flag
		friend.CreateTime = time.Now()
		friendReq.CreateTime = time.Now()
		err = model.UpdateFriendRequest(&friendReq)
		if err == errorcode.DataErr {
			code, msg := errorcode.ParserErr(err)
			c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		}
		err = model.SaveFriend(&friend)
		code, msg := errorcode.ParserErr(err)
		user, _ := model.UserInfoByUid(uid)
		var tmpf model.FriendInfo
		tmpf.Uid = user.Uid
		tmpf.Remark = ""
		tmpf.NickName = user.NickName
		tmpf.Phone = user.Phone
		tmpf.Email = user.Email
		tmpf.Ex = user.Ex
		tmpf.Gender = user.Gender
		tmpf.Birth = user.Birth
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		if code == errorcode.Ok.Code {
			var tmp model.MsgData
			tmp.Platform = ""
			tmp.SendId = friendReq.UserId
			tmp.ReceiveId = friendReq.ReqId
			tmp.ContentType = 0
			tmp.Content = "已经添加你为好友啦，开始聊天吧~"
			msgData, _ := model.SaveMsg(&tmp)
			res := make(map[string]interface{})
			res["msg"] = &msgData
			res["friend"] = &tmpf
			if msgData != nil && socket_io.SocketHub.SocketClients[friendReq.ReqId] != nil {
				socket_io.Server.BroadcastToRoom("/", friendReq.ReqId, "receive_due_friend", res)
			}
		}
		break
	case enum.Refuse.Flag:
		friendReq.Flag = enum.Refuse.Flag
		err = model.DeleteFriendReq(&friendReq)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		break
	case enum.Ignore.Flag:
		friendReq.Flag = enum.Refuse.Flag
		err = model.DeleteFriendReq(&friendReq)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		break
	default:
	}
	return
}

func GetAllFriendReqs(c *gin.Context) {
	var token auth.Token
	err := c.ShouldBindJSON(&token)
	if err != nil {
		fmt.Println(err)
		return
	}
	uid, tokenErr := auth.ValidateToken(token.Token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	friendReqsR, friendReqsRErr := model.QueryFriendRequestsReceive(uid)
	//friendReqsS, friendReqsSErr := model.QueryFriendRequestsSend(uid)
	//if friendReqsRErr == errorcode.DataErr || friendReqsSErr == errorcode.DataErr {
	//	c.JSON(http.StatusOK, gin.H{"code": errorcode.DataErr.Code, "message": errorcode.DataErr.Message})
	//	return
	//}
	if friendReqsRErr == errorcode.DataErr {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.DataErr.Code, "message": errorcode.DataErr.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "req_receive": friendReqsR})
	return
}

func GetAllFriends(c *gin.Context) {
	var token auth.Token
	err := c.ShouldBindJSON(&token)
	if err != nil {
		fmt.Println(err)
		return
	}
	uid, tokenErr := auth.ValidateToken(token.Token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	friends, friendErr := model.QueryFriends(uid)
	if friendErr != nil {
		code, msg := errorcode.ParserErr(friendErr)
		if code == errorcode.FriendsNo.Code {
			c.JSON(http.StatusOK, gin.H{"code": code, "message": msg, "friends": ""})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		return
	}
	var friendInfos []model.FriendInfo
	var friendInfo model.FriendInfo
	for _, friend := range friends {
		user, err := model.UserInfoByUid(friend.FriendId)
		if err != nil {
			code, msg := errorcode.ParserErr(friendErr)
			c.JSON(http.StatusOK, gin.H{"code": code, "message": msg, "friends": ""})
			return
		}
		friendInfo.Remark = friend.Comment
		friendInfo.Uid = user.Uid
		friendInfo.NickName = user.NickName
		friendInfo.Phone = user.Phone
		friendInfo.Email = user.Email
		friendInfo.Birth = user.Birth
		friendInfo.HeadImg = user.HeadImg
		friendInfo.Gender = user.Gender
		friendInfo.Ex = user.Ex
		friendInfos = append(friendInfos, friendInfo)
	}
	c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "friends": friendInfos})
	return
}

func UpdateFriendComment(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var friend model.Friend
	c.ShouldBindJSON(&friend)
	err := model.UpdateFriendComment(&friend)
	code, msg := errorcode.ParserErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	return
}

func DeleteFriend(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var friend model.Friend
	c.ShouldBindJSON(&friend)
	err := model.UpdateMsgDelFlag(friend.OwnerId, friend.FriendId)
	if err != nil {
		err = model.DeleteFriend(&friend)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		return
	}
	isFriend := model.CheckIsFriend(&friend)
	if isFriend == "true" {
		err = model.DeleteFriend(&friend)
		code, msg := errorcode.ParserErr(err)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message})
	}
	return
}
