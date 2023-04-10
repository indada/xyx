package routes

import (
	"github.com/gin-gonic/gin"
	"xyx/controller"
)

func OpenApi() *gin.Engine {
	r := gin.Default()

	//接收抖音头条数据用户组
	dyGroup := r.Group("dy")
	dyGroup.POST("/discuss", controller.CallbackComment) //评论
	dyGroup.POST("/gift", controller.CallbackGift)       //礼物
	dyGroup.POST("/likes", controller.CallbackLike)      //点赞
	dyGroup.POST("/css", controller.CallbackCss)         //cs
	//接收小玩法请求数据用户组
	xwfGroup := r.Group("xwf")
	xwfGroup.POST("/firing", controller.Firing)         //开启小玩法
	xwfGroup.POST("/cease", controller.Cease)           //关闭
	xwfGroup.POST("/getMessage", controller.GetMessage) //获取数据
	xwfGroup.POST("/addMessage", controller.AddMessage) //关闭

	xwfSocket := r.Group("socket")
	xwfSocket.Any("/getMessage", controller.GetMessageSocket) //长连接

	return r
}
