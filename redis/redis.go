package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

var rd *redis.Client

func init() {
	// 全局变量：连接数据库
	rd = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // url
		Password: "123456",
		DB:       0, // 数据库
	})
	result, err := rd.Ping().Result()
	if err != nil {
		fmt.Println("redis 连接失败 ping err :", err)
		return
	}
	fmt.Println(result)
}

func ConnRedis() { //测试连接
	rd := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // url
		Password: "123456",
		DB:       0, // 数据库
	})
	result, err := rd.Ping().Result()
	if err != nil {
		fmt.Println("redis 连接失败 ping err :", err)
		return
	}
	fmt.Println(result)
}

var NN = redis.Nil

func SetCache(key string, val string, exp time.Duration) error {
	err := rd.Set(key, val, exp).Err()
	return err
}
func GetCache(key string) (string, error) {
	val, err := rd.Get(key).Result()
	return val, err
}
func AddPush(key string, val interface{}) {
	rd.RPush(key, val)
}
func GetPop(key string) (string, error) {
	first, err := rd.LPop(key).Result()
	return first, err
}
func GetBLPop(key string) (string, error) {
	first, err := rd.BLPop(time.Second*60, key).Result()
	if err != nil {
		return "", err
	}
	return first[1], err
}
func GetListLen(key string) (int64, error) {
	return rd.LLen(key).Result()
}
func GetRangeAll(key string) []string {
	list, _ := rd.LRange(key, 0, -1).Result()
	return list
}
func HmSet(key string, fields map[string]interface{}) {
	rd.HMSet(key, fields)
}
func HSet(key, field, val string) {
	rd.HSet(key, field, val)
}
func HGet(key string, field string) (string, error) {
	first, err := rd.HGet(key, field).Result()
	return first, err
}
func GetKey(s1, s2 string) string {
	var bd strings.Builder
	bd.WriteString(s1)
	bd.WriteString("_")
	bd.WriteString(s2)
	return bd.String()
}
func GetKeys(s1, s2, s3 string) string {
	var bd strings.Builder
	bd.WriteString(s1)
	bd.WriteString("_")
	bd.WriteString(s2)
	bd.WriteString("_")
	bd.WriteString(s3)
	return bd.String()
}
