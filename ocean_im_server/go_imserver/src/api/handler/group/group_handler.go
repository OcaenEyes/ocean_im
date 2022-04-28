package group

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"oceanim/go_server/src/api/enum"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/model"
	"oceanim/go_server/src/utils/auth"
	"time"
)

func CreateGroupHandler(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var groupTmp model.GroupTmp
	var group model.Group

	paramErr := c.ShouldBindJSON(groupTmp)
	if paramErr != nil {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupReqParamErr.Code, "message": errorcode.GroupReqParamErr.Message})
		return
	}
	var groupExist bool
	var groupId string
	for groupExist {
		groupId = fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
		groupExist = model.CheckGroupIsExit(groupId)
	}
	group.GroupId = groupId
	group.GroupName = groupTmp.GroupName
	group.FaceImg = groupTmp.FaceImg
	group.Notification = groupTmp.Notification
	group.Introduction = groupTmp.Introduction
	group.Ex = groupTmp.Ex
	group.CreateTime = time.Now()
	saveRes := model.SaveGroup(&group)

	groupAdmin := model.GroupMember{
		GroupId:            groupId,
		UserId:             uid,
		NickName:           "",
		AdministratorLevel: enum.Creator.Level,
		JoinTime:           time.Now(),
	}
	saveRes = model.JoinGroup(&groupAdmin)

	for _, id := range groupTmp.MemberUid {
		groupMember := model.GroupMember{
			GroupId:            groupId,
			UserId:             id,
			NickName:           "",
			AdministratorLevel: enum.Creator.Level,
			JoinTime:           time.Now(),
		}
		saveRes = model.JoinGroup(&groupMember)
	}
	code, msg := errorcode.ParserErr(saveRes)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	return
}

func LeaveGroupHandler(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var member model.GroupMember
	paramErr := c.ShouldBindJSON(&member)
	if paramErr != nil || member.UserId == "" || member.GroupId == "" {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupReqParamErr.Code, "message": errorcode.GroupReqParamErr.Message})
		return
	}
	leaveRes := model.DeleteGroupMember(member.UserId, member.GroupId)
	code, msg := errorcode.ParserErr(leaveRes)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	return
}

func RemoveGroupMemberHandler(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, tokenErr := auth.ValidateToken(token)
	if tokenErr != nil || len(uid) == 0 {
		fmt.Println(tokenErr)
		fmt.Println(uid)
		c.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
		return
	}
	var member model.GroupMember
	paramErr := c.ShouldBindJSON(&member)
	if paramErr != nil || member.UserId == "" || member.GroupId == "" {
		c.JSON(http.StatusOK, gin.H{"code": errorcode.GroupReqParamErr.Code, "message": errorcode.GroupReqParamErr.Message})
		return
	}
	leaveRes := model.DeleteGroupMember(member.UserId, member.GroupId)
	code, msg := errorcode.ParserErr(leaveRes)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": msg})
	return
}
