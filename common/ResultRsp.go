package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultRsp struct {
	code int         `json:"code"`
	msg  string      `json:"success"`
	data interface{} `json:data"`
}

// 正确状态处理
func Success(c *gin.Context, msg string, data interface{}) {
	if data == "" {
		data = struct{}{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

// 错误状态处理
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  msg,
	})
}
