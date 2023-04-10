package dao

import orm "xyx/databases"

type TbDy struct {
	Id         int64  `gorm:"column:id;primary_key;auto_increment"`
	Nickname   string `gorm:"column:nickname;type:varchar(255);not null"`
	SecOpenid  string `gorm:"column:sec_openid;type:varchar(255);not null"`
	AvatarUrl  string `gorm:"column:avatar_url;type:varchar(1000);not null"`
	SecGiftId  string `gorm:"column:sec_gift_id;type:varchar(255);not null"`
	Content    string `gorm:"column:content;type:varchar(255);not null"`
	LikeNum    string `gorm:"column:like_num;type:varchar(255);not null"`
	Roomid     string `gorm:"column:roomid;type:varchar(255);not null"`
	GiftNum    int64  `gorm:"column:gift_num;type:int(10);not null"`
	GiftValue  int64  `gorm:"column:sec_gift_id;type:bigint(11);not null"`
	MsgType    int    `gorm:"column:msg_type;type:tinyint(4);not null"`
	CreateTime int64  `gorm:"column:create_time;int(10);not null"`
	UserId     int64  `gorm:"column:user_id;int(10);not null"`
}

func (t *TbDy) TableName() string {
	return "tb_dy"
}
func CreateDy(data TbDy) (int64, error) {
	db := orm.Eloquent
	result := db.Create(&data)
	return data.Id, result.Error
}
