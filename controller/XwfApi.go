package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"time"
	"xyx/common"
	"xyx/config"
	"xyx/entity"
	"xyx/service"
)

// 开启
func Firing(c *gin.Context) {
	var dd entity.FringData
	c.BindJSON(&dd)
	fmt.Println("dd:", dd)
	dd.Appid = config.Appid
	if dd.MsgType == "all" {
		g, _ := errgroup.WithContext(context.Background())
		g.Go(func() error {
			dd.MsgType = "live_comment"
			time.Sleep(time.Second) //优化空间
			return service.FiringGame(dd)
		})
		g.Go(func() error {
			dd.MsgType = "live_gift"
			//time.Sleep(3 * time.Second)
			return service.FiringGame(dd)
		})
		g.Go(func() error {
			dd.MsgType = "live_like"
			return service.FiringGame(dd)
		})

		err := g.Wait()
		if err != nil {
			common.Error(c, err.Error())
		} else {
			common.Success(c, "开启成功", "")
		}
	} else {
		err := service.FiringGame(dd)
		if err != nil {
			common.Error(c, err.Error())
		} else {
			common.Success(c, "开启成功", "")
		}
	}

}

// 关闭
func Cease(c *gin.Context) {
	var dd entity.FringData
	c.BindJSON(&dd)
	dd.Appid = config.Appid
	if dd.MsgType == "all" {
		g, _ := errgroup.WithContext(context.Background())
		g.Go(func() error {
			dd.MsgType = config.LiveComment
			time.Sleep(time.Second) //优化空间
			return service.CeaseGame(dd)
		})
		g.Go(func() error {
			dd.MsgType = config.LiveGift
			//time.Sleep(3 * time.Second)
			return service.CeaseGame(dd)
		})
		g.Go(func() error {
			dd.MsgType = config.LiveLike
			return service.CeaseGame(dd)
		})

		err := g.Wait()
		if err != nil {
			common.Error(c, err.Error())
		} else {
			common.Success(c, "关闭成功", "")
		}
	} else {
		err := service.CeaseGame(dd)
		if err != nil {
			common.Error(c, err.Error())
		} else {
			common.Success(c, "关闭成功", "")
		}
	}
}

// 获取数据
func GetMessage(c *gin.Context) {
	var dd entity.FringData
	c.BindJSON(&dd)
	dataBack := make(map[string]any)
	if dd.MsgType == "all" {
		dd.MsgType = config.LiveComment
		dataBack[dd.MsgType] = service.CallbackXyx(dd)
		dd.MsgType = config.LiveGift
		dataBack[dd.MsgType] = service.CallbackXyx(dd)
		dd.MsgType = config.LiveLike
		dataBack[dd.MsgType] = service.CallbackXyx(dd)
		common.Success(c, "获取成功!", dataBack)
	} else {
		daa := service.CallbackXyx(dd)
		if len(daa) == 0 {
			common.Success(c, "获取成功!", struct{}{})
		} else {
			dataBack[dd.MsgType] = daa
			common.Success(c, "获取成功!", dataBack)
		}

	}

}

// 添加数据
func AddMessage(c *gin.Context) {
	var dd entity.FringData
	c.BindJSON(&dd)
	dd.MsgType = config.LiveComment
	for i := 0; i < 50; i++ {
		add := entity.Msg{}
		add.AvatarUrl = "https://dd.com/12.png"
		add.MsgId = common.RandStr(6)
		add.Roomid = dd.Roomid
		add.MsgType = config.LiveComment
		add.SecOpenid = common.RandStr(6)
		add.Content = common.RandStr(6)
		add.Nickname = common.RandStr(4)
		add.Timestamp = time.Now().UnixMicro()
		service.TestAdd(dd, add)
	}
	dd.MsgType = config.LiveGift
	for i := 0; i < 50; i++ {
		add := entity.Msg{}
		add.AvatarUrl = "https://dd.com/12.png"
		add.MsgId = common.RandStr(6)
		add.Roomid = dd.Roomid
		add.MsgType = config.LiveGift
		add.SecOpenid = common.RandStr(6)
		add.SecGiftId = common.RandStr(6)
		add.GiftNum = 2
		add.GiftValue = 2000
		add.Nickname = common.RandStr(4)
		add.Timestamp = time.Now().UnixMicro()
		service.TestAdd(dd, add)
	}
	dd.MsgType = config.LiveLike
	for i := 0; i < 50; i++ {
		add := entity.Msg{}
		add.AvatarUrl = "https://dd.com/12.png"
		add.Roomid = dd.Roomid
		add.MsgType = config.LiveLike
		add.MsgId = common.RandStr(6)
		add.SecOpenid = common.RandStr(6)
		add.LikeNum = "2"
		add.Nickname = common.RandStr(4)
		add.Timestamp = time.Now().UnixMicro()
		service.TestAdd(dd, add)
	}
	common.Success(c, "成功!", struct{}{})
}
