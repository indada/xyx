package entity

type FringData struct {
	Roomid  string `json:"roomid"`
	Appid   string `json:"appid"`
	MsgType string `json:"msg_type"`
}
type GetAccessToken struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
}
type Access struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
	Data    struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
}
type TaskStartBack struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Logid  string `json:"logid"`
	Data   struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}

type Msg struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	Content   string `json:"content"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
	SecGiftId string `json:"sec_gift_id"`
	GiftNum   int64  `json:"gift_num"`
	GiftValue int32  `json:"gift_value"`
	LikeNum   string `json:"like_num"`
	Roomid    string `json:"roomid"`
	MsgType   string `json:"msg_type"`
}
type Gift struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	SecGiftId string `json:"sec_gift_id"`
	GiftNum   int    `json:"gift_num"`
	GiftValue int    `json:"gift_value"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}
type Like struct {
	MsgId     string `json:"msg_id"`
	SecOpenid string `json:"sec_openid"`
	LikeNum   string `json:"like_num"`
	AvatarUrl string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	Timestamp int64  `json:"timestamp"`
}
