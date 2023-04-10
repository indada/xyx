package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"xyx/config"
	"xyx/redis"
)

var upGrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

var Clients = make(map[string]*websocket.Conn)

func GetMessageSocket(c *gin.Context) {
	Roomid := c.Query("roomid") //var ws = new WebSocket("ws://localhost:2303/client?user="+user); 直播间id
	ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
	Clients[Roomid] = ws
	done := make(chan struct{})
	defer ws.Close()
	//发数据
	msg := make(chan string)
	go func() {
		listKey := redis.GetKeys(config.ListKey, Roomid, config.LiveComment)
		fmt.Println("get_redis_Key", listKey)
		for {
			_, ok := Clients[Roomid]
			if !ok {
				break
			}
			datum, err := redis.GetPop(listKey)
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}
			msg <- datum
		}

	}()
	go func() {
		listKey := redis.GetKeys(config.ListKey, Roomid, config.LiveGift)
		fmt.Println("get_redis_Key", listKey)
		for {
			_, ok := Clients[Roomid]
			if !ok {
				break
			}
			datum, err := redis.GetPop(listKey)
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}
			msg <- datum
		}

	}()
	go func() {
		listKey := redis.GetKeys(config.ListKey, Roomid, config.LiveLike)
		fmt.Println("get_redis_Key", listKey)
		for {
			_, ok := Clients[Roomid]
			if !ok {
				break
			}
			datum, err := redis.GetPop(listKey)
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			}
			msg <- datum
		}

	}()

	go func() {
	LOOP:
		for {
			select {
			case datum := <-msg:
				var da map[string]interface{}
				json.Unmarshal([]byte(datum), &da)
				ws.WriteJSON(da)
			case <-done: //退出了
				break LOOP
			}
		}

	}()

	for {
		mt, _, err := ws.ReadMessage()
		if err != nil { //关闭了
			delete(Clients, Roomid)
			done <- struct{}{}
			fmt.Println("ReadMessage err:", err)
			break
		}
		fmt.Println("ReadMessage mt:", mt)
		//msg <- string(message)
		/*err = ws.WriteMessage(mt, message)

		if err != nil {
			break
		}*/
	}

}
