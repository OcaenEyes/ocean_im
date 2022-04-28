package routers

import (
	"github.com/gin-gonic/gin"
	"oceanim/go_server/src/api/handler/message"
)

func MsgRouter(e *gin.Engine) {
	msgHandler := e.Group("msg")
	msgHandler.POST("v1/getAllMsg", message.QueryAllMsgHandler)
}
