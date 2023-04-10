package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xyx/common"
	"xyx/service"
)

// 评论回调
func CallbackComment(c *gin.Context) {
	header := map[string]string{
		"x-nonce-str": c.Request.Header.Get("x-nonce-str"),
		"x-timestamp": c.Request.Header.Get("x-timestamp"),
		"x-roomid":    c.Request.Header.Get("x-roomid"),
		"x-msg-type":  c.Request.Header.Get("x-msg-type"),
	}
	signature := common.Signature(header)
	if signature != c.Request.Header.Get("x-signature") {
		common.Error(c, "密钥错误!")
	} else {
		var dd []map[string]interface{}
		c.BindJSON(&dd)
		err := service.CallbackDy(header, dd)
		fmt.Println(err)
		common.Success(c, "开启成功", "")
	}

}

// 礼物回调
func CallbackGift(c *gin.Context) {
	header := map[string]string{
		"x-nonce-str": c.Request.Header.Get("x-nonce-str"),
		"x-timestamp": c.Request.Header.Get("x-timestamp"),
		"x-roomid":    c.Request.Header.Get("x-roomid"),
		"x-msg-type":  c.Request.Header.Get("x-msg-type"),
	}
	signature := common.Signature(header)
	if signature != c.Request.Header.Get("x-signature") {
		common.Error(c, "密钥错误!")
	} else {
		var dd []map[string]interface{}
		c.BindJSON(&dd)
		err := service.CallbackDy(header, dd)
		fmt.Println(err)
		common.Success(c, "开启成功", "")
	}

}

// 点赞回调
func CallbackLike(c *gin.Context) {
	header := map[string]string{
		"x-nonce-str": c.Request.Header.Get("x-nonce-str"),
		"x-timestamp": c.Request.Header.Get("x-timestamp"),
		"x-roomid":    c.Request.Header.Get("x-roomid"),
		"x-msg-type":  c.Request.Header.Get("x-msg-type"),
	}
	signature := common.Signature(header)
	if signature != c.Request.Header.Get("x-signature") {
		common.Error(c, "密钥错误!")
	} else {
		var dd []map[string]interface{}
		c.BindJSON(&dd)
		err := service.CallbackDy(header, dd)
		fmt.Println(err)
		common.Success(c, "开启成功", "")
	}

}

// 测试回调
func CallbackCss(c *gin.Context) {
	header := map[string]string{
		"x-nonce-str": c.Request.Header.Get("x-nonce-str"),
		"x-timestamp": c.Request.Header.Get("x-timestamp"),
		"x-roomid":    c.Request.Header.Get("x-roomid"),
		"x-msg-type":  c.Request.Header.Get("x-msg-type"),
	}
	signature := common.Signature(header)
	fmt.Println("signature:", signature)
	if signature != c.Request.Header.Get("x-signature") {
		common.Error(c, "密钥错误!")
	} else {
		var dd []map[string]interface{}
		c.BindJSON(&dd)
		err := service.CallbackDy(header, dd)
		fmt.Println(err)
		common.Success(c, "开启成功", "")
	}

}
