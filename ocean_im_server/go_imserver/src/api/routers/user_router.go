package routers

import (
	"github.com/gin-gonic/gin"
	"oceanim/go_server/src/api/handler/user"
)

func UserRouter(e *gin.Engine) {
	userHandler := e.Group("user")
	userHandler.POST("v1/register", user.Register)
	userHandler.POST("v1/login", user.Login)
	userHandler.POST("v1/getUserInfo", user.GetUserInfo)
	userHandler.POST("v1/getUserInfoByUid", user.GetUserInfoByUid)
	userHandler.POST("v1/getUserInfoByPhone", user.GetUserInfoByPhone)
	userHandler.POST("v1/updateUserInfo", user.UpdateUserInfo)
}
