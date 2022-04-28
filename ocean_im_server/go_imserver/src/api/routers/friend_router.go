package routers

import (
	"github.com/gin-gonic/gin"
	"oceanim/go_server/src/api/handler/friend"
)

func FriendRouters(e *gin.Engine) {
	friendHandler := e.Group("friend")
	friendHandler.POST("v1/getAllFriends", friend.GetAllFriends)
	friendHandler.POST("v1/getAllFriendReqs", friend.GetAllFriendReqs)
	friendHandler.POST("v1/sendFriendReq", friend.SendFriendReq)
	friendHandler.POST("v1/dueFriendReq", friend.DueFriendReq)
	friendHandler.POST("v1/deleteFriend", friend.DeleteFriend)
	friendHandler.POST("v1/updateFriendRemark", friend.UpdateFriendComment)

}
