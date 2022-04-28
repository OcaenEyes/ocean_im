package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanim/go_server/src/api/errorcode"
	"oceanim/go_server/src/api/routers"
	"oceanim/go_server/src/common/db"
	"oceanim/go_server/src/utils/auth"
	"oceanim/go_server/src/utils/cors"
	"oceanim/go_server/src/utils/socket_io"
)

//var hub = ws.NewMyWsHub()



func init() {
	db.InitMysqlTable()
	//go hub.Run()

}
func main() {
	r := gin.Default()
	r.Use(cors.MyCorsHandler())
	routers.UserRouter(r)
	routers.FriendRouters(r)
	routers.MsgRouter(r)
	//r.GET("ws", func(c *gin.Context) {
	//	ws.MyWsServer(hub, c)
	//})
	go func() {
		err := socket_io.Server.Serve()
		if err != nil {
			fmt.Println(err)
		}
	}()
	defer socket_io.Server.Close()
	r.GET("/ws", func(context *gin.Context) {
		context.Request.Header.Del("Origin")
		token := context.Query("token")
		_, tokenErr := auth.ValidateToken(token)
		if tokenErr != nil {
			fmt.Println(tokenErr)
			context.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
			return
		}
		socket_io.Server.ServeHTTP(context.Writer, context.Request)
	})
	r.POST("/ws", func(context *gin.Context) {
		context.Request.Header.Del("Origin")
		token := context.Query("token")
		_, tokenErr := auth.ValidateToken(token)
		if tokenErr != nil {
			fmt.Println(tokenErr)
			context.JSON(http.StatusOK, gin.H{"code": errorcode.AuthErr.Code, "message": errorcode.AuthErr.Message})
			return
		}
		socket_io.Server.ServeHTTP(context.Writer, context.Request)
	})
	r.Run()
}
