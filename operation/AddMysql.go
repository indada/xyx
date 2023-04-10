package operation

import (
	"encoding/json"
	"fmt"
	"time"
	"xyx/config"
	"xyx/dao"
	"xyx/redis"
)

func GetList() {
	listAddKey := config.ListKey //用户保存至Mysql数据库
	ch := make(chan int, 50)     //创建大小为50的缓冲通道

	for {
		datum, err := redis.GetBLPop(listAddKey)
		if err != nil {
			fmt.Println("redis err:", err)
			time.Sleep(time.Second * 1)
			continue
		}
		ch <- 1 //大于50时就会堵塞
		go addMysql(datum, ch)

	}
}
func addMysql(datum string, ch chan int) {
	var da map[string]interface{}
	json.Unmarshal([]byte(datum), &da)
	//time.Sleep(time.Second * 2)
	//fmt.Println("addMysql:", da)
	defer func() {
		<-ch
	}()
	addUser(da)
}
func addUser(da map[string]interface{}) {
	users, err := dao.GetUserByOpenId(da["sec_openid"].(string))
	if err != nil {
		//为空 可以添加
		users = dao.TbUser{
			Nickname:   da["nickname"].(string),
			SecOpenid:  da["sec_openid"].(string),
			AvatarUrl:  da["avatar_url"].(string),
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}
		userid, err := dao.CreateUser(users)
		if err != nil {
			fmt.Println("添加用户数据失败!", err)
			return
		}
		users.Id = userid
	}
	msgType := da["msg_type"].(string)
	fmt.Println("da:", da)
	var dy dao.TbDy
	if msgType == "live_gift" { //礼物
		//添加数据
		dy = dao.TbDy{
			Nickname:   da["nickname"].(string),
			SecOpenid:  da["sec_openid"].(string),
			AvatarUrl:  da["avatar_url"].(string),
			SecGiftId:  da["sec_gift_id"].(string),
			Roomid:     da["roomid"].(string),
			GiftNum:    int64(da["gift_num"].(float64)),
			GiftValue:  int64(da["gift_value"].(float64)),
			MsgType:    1,
			CreateTime: time.Now().Unix(),
			UserId:     users.Id,
		}
	} else if msgType == "live_like" { //点赞
		dy = dao.TbDy{
			Nickname:   da["nickname"].(string),
			SecOpenid:  da["sec_openid"].(string),
			AvatarUrl:  da["avatar_url"].(string),
			Roomid:     da["roomid"].(string),
			LikeNum:    da["like_num"].(string),
			MsgType:    2,
			CreateTime: time.Now().Unix(),
			UserId:     users.Id,
		}
	} else { //评论
		dy = dao.TbDy{
			Nickname:   da["nickname"].(string),
			SecOpenid:  da["sec_openid"].(string),
			AvatarUrl:  da["avatar_url"].(string),
			Roomid:     da["roomid"].(string),
			Content:    da["content"].(string),
			MsgType:    0,
			CreateTime: time.Now().Unix(),
			UserId:     users.Id,
		}
	}
	dao.CreateDy(dy)
	//fmt.Println("userId:", users.Id, "dyId:", dyId, "err", err)
}
