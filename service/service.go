package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"xyx/config"
	"xyx/entity"
	"xyx/redis"
)

// 开启游戏
func FiringGame(dd entity.FringData) error {
	token, err := GetAccessToken() //获取token
	if err != nil {
		return err
	}
	//开启任务
	fmt.Println("开启dd:", dd, "token:", token)
	result, err := postRequest(config.TastStartUrlTest, dd)
	if err != nil {
		return err
	}
	req := entity.TaskStartBack{}
	err = json.Unmarshal(result, &req)
	if err != nil {
		return err
	}
	if req.ErrNo != 0 {
		return errors.New(req.ErrMsg)
	}
	return nil
}

// 关闭游戏
func CeaseGame(dd entity.FringData) error {
	token, err := GetAccessToken() //获取token
	if err != nil {
		return err
	}
	//开启任务
	fmt.Println("关闭dd:", dd, "token:", token)
	result, err := postRequest(config.TastCeasetUrlTest, dd)
	if err != nil {
		return err
	}
	req := entity.TaskStartBack{}
	err = json.Unmarshal(result, &req)
	if err != nil {
		return err
	}
	if req.ErrNo != 0 {
		return errors.New(req.ErrMsg)
	}
	return nil
}

// 接收抖音数据
func CallbackDy(header map[string]string, body []map[string]interface{}) error {
	hashKey := redis.GetKey(config.HashKey, header["x-roomid"])
	listAddKey := config.ListKey //用户保存至Mysql数据库
	listKey := redis.GetKeys(config.ListKey, header["x-roomid"], header["x-msg-type"])

	for _, msg := range body {
		msgId, _ := msg["msg_id"]
		first, err := redis.HGet(hashKey, msgId.(string))
		if first == "" || err != nil {
			redis.HGet(hashKey, "ok")
			msg["roomid"] = header["x-roomid"]
			msg["msg_type"] = header["x-msg-type"]
			msgJson, _ := json.Marshal(msg)
			strMsg := string(msgJson)
			redis.AddPush(listKey, strMsg)
			redis.AddPush(listAddKey, strMsg)
		}
	}
	return nil
}
func TestAdd(dd entity.FringData, msg entity.Msg) {
	listKey := redis.GetKeys(config.ListKey, dd.Roomid, dd.MsgType)
	fmt.Println("add_redis_Key", listKey)
	listAddKey := config.ListKey //用户保存至Mysql数据库
	msgJson, _ := json.Marshal(msg)
	strMsg := string(msgJson)
	redis.AddPush(listKey, strMsg)
	redis.AddPush(listAddKey, strMsg)
}

func CallbackXyx(dd entity.FringData) []map[string]interface{} {
	var backData []map[string]interface{}
	listKey := redis.GetKeys(config.ListKey, dd.Roomid, dd.MsgType)
	a, err := redis.GetListLen(listKey)
	if err != nil {
		return backData
	}
	if a == 0 {
		return backData
	}
	var i int64
	for i = 0; i <= a; i++ {
		datum, err := redis.GetPop(listKey)
		if err != nil {
			return backData
		}
		if datum == "" {
			return backData
		}
		var da map[string]interface{}
		json.Unmarshal([]byte(datum), &da)
		backData = append(backData, da)
	}
	/*allData := redis.GetRangeAll(listKey)
	for _, datum := range allData {
		var da map[string]interface{}
		json.Unmarshal([]byte(datum), &da)
		backData = append(backData, da)
	}*/
	return backData
}

// 获取access_token
func GetAccessToken() (string, error) {
	val, err := redis.GetCache(config.TokenKey)
	if err != redis.NN && val != "" {
		return val, nil
	}
	getDd := entity.GetAccessToken{
		config.Appid,
		config.Secret,
		config.GrantType,
	}
	result, err := postRequest(config.AtUrlGTest, getDd)
	if err != nil {
		return "", err
	}
	req := entity.Access{}
	err = json.Unmarshal(result, &req)
	if err != nil {
		return "", err
	}
	if req.ErrNo != 0 {
		return "", errors.New(req.ErrTips)
	}
	//保存token
	err = redis.SetCache(config.TokenKey, req.Data.AccessToken, time.Hour)
	return req.Data.AccessToken, err
}

// post请求
func postRequest(url string, data interface{}) ([]byte, error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	return result, err
}
