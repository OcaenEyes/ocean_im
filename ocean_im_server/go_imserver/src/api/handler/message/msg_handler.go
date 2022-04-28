package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/model"
	"oceanim/go_server/src/utils/auth"
)

func QueryAllMsgHandler(c *gin.Context) {
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

	msgAll, msgErr := model.SelectMsg(uid)

	if msgErr != nil {
		code, errMsg := errorcode.ParserErr(msgErr)
		if code == errorcode.MsgNo.Code {
			c.JSON(http.StatusOK, gin.H{"code": code, "message": errMsg, "chats": ""})
			return

		}
		c.JSON(http.StatusOK, gin.H{"code": code, "message": errMsg})
		return
	}
	msgS := make(map[string][]model.ChatMsg)
	for _, msg := range msgAll {
		if msg.SendId == uid {
			_, ok := msgS[msg.ReceiveId]
			if ok {
				msgS[msg.ReceiveId] = append(msgS[msg.ReceiveId], msg)
			} else {
				var tmpS []model.ChatMsg
				tmpS = append(tmpS, msg)
				msgS[msg.ReceiveId] = tmpS
			}
		} else {
			_, ok := msgS[msg.SendId]
			if ok {
				msgS[msg.SendId] = append(msgS[msg.SendId], msg)
			} else {
				var tmpS []model.ChatMsg
				tmpS = append(tmpS, msg)
				msgS[msg.SendId] = tmpS
			}
		}
	}
	fmt.Println(msgS)

	var resMsg []interface{}
	for k,v := range msgS {
		temp := make(map[string]interface{})
		temp["to_user"]=k
		temp["detail_list"]=v
		resMsg = append(resMsg, temp)
	}
	c.JSON(http.StatusOK, gin.H{"code": errorcode.Ok.Code, "message": errorcode.Ok.Message, "chats": resMsg})
	return
}
