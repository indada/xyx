package config

// url
var AtUrlGTest string = "http://dc.com/index.php"
var TastStartUrlTest string = "http://dc.com/index.php"
var TastCeasetUrlTest string = "http://dc.com/index.php"

var AtUrl string = "https://developer.toutiao.com/api/apps/v2/token"               //获取token
var TastStartUrl string = "https://webcast.bytedance.com/api/live_data/task/start" //开启游戏
var TastCeasetUrl string = "https://webcast.bytedance.com/api/live_data/task/stop" //关闭游戏

// 密钥等数据
var BodyStr string = "abc123"
var Secret string = "123"
var Appid string = "123"

var GrantType string = "client_credential"

var LiveComment string = "live_comment" //评论
var LiveGift string = "live_gift"       //礼物
var LiveLike string = "live_like"       //点赞

// redisKey
var TokenKey string = "DYACCESSTOKEN"
var HashKey string = "HashKey"
var ListKey string = "ListKey"

/**
func (d Duration) Nanoseconds() int64 {}   // 纳秒
func (d Duration) Microseconds() int64 {}  // 微秒
func (d Duration) Milliseconds() int64 {}  // 毫秒
func (d Duration) Seconds() float64 {}     // 秒
func (d Duration) Minutes() float64 {}     // 分钟
func (d Duration) Hours() float64 {}       // 小时
*/
