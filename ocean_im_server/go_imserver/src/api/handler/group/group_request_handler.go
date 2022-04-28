package group

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanim/go_server/src/api/enum"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/model"
	"oceanim/go_server/src/utils/auth"
)

func SendJoinGroupReqHandler(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var groupReq model.GroupRequest
	paramErr := c.ShouldBindJSON(&groupReq)
	if paramErr != nil || groupReq.GroupId == "" || groupReq.FromUserId == "" {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupReqParamErr.Code, "message": errorcode.GroupReqParamErr.Message})
		return
	}

	groupExist := model.CheckGroupIsExit(groupReq.GroupId)
	if !groupExist {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupsNo.Code, "message": errorcode.GroupsNo.Message})
		return
	}

	saveRes := model.SaveGroupRequest(&groupReq)
	code, msg := errorcode.ParserErr(saveRes)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	return
}

func DueJoinGroupReqHandler(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var groupReq model.GroupRequest
	paramErr := c.ShouldBindJSON(&groupReq)
	if paramErr != nil || groupReq.GroupId == "" || groupReq.FromUserId == "" || groupReq.HandledUser == "" {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupReqParamErr.Code, "message": errorcode.GroupReqParamErr.Message})
		return
	}
	groupExist := model.CheckGroupIsExit(groupReq.GroupId)
	if !groupExist {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupsNo.Code, "message": errorcode.GroupsNo.Message})
		return
	}

	switch groupReq.Flag {
	case enum.AcceptJoin.Status:
		saveRes := model.UpdateGroupRequest(&groupReq)
		code, msg := errorcode.ParserErr(saveRes)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		break
	case enum.IgnoreJoin.Status:
		saveRes := model.DeleteGroupRequest(&groupReq)
		code, msg := errorcode.ParserErr(saveRes)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		break
	case enum.RefuseJoin.Status:
		saveRes := model.DeleteGroupRequest(&groupReq)
		code, msg := errorcode.ParserErr(saveRes)
		c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
		break
	default:
	}
	return
}
