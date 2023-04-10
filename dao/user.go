package dao

import orm "xyx/databases"

type TbUser struct {
	Id         int64  `gorm:"column:id;primary_key;auto_increment"`
	Nickname   string `gorm:"column:nickname;type:varchar(255);not null"`
	SecOpenid  string `gorm:"column:sec_openid;type:varchar(255);not null"`
	AvatarUrl  string `gorm:"column:avatar_url;type:varchar(1000);not null"`
	CreateTime int64  `gorm:"column:create_time;int(10);not null"`
	UpdateTime int64  `gorm:"column:update_time;int(10);not null"`
}

func GetTbUser() TbUser {
	return TbUser{}
}
func (t *TbUser) TableName() string {
	return "tb_user"
}

func GetUserByOpenId(openId string) (TbUser, error) {
	model := TbUser{}
	db := orm.Eloquent
	result := db.Where("sec_openid = ?", openId).First(&model)
	return model, result.Error
}
func CreateUser(data TbUser) (int64, error) {
	db := orm.Eloquent
	result := db.Create(&data)
	return data.Id, result.Error
}

//------------------------

func ListTest(datas []TbUser, page int) ([]TbUser, int64, error) {
	var pageSize = 2
	db := orm.Eloquent
	offset := (page - 1) * pageSize
	result := db.Order("id desc").Offset(offset).Limit(pageSize).Find(&datas)
	return datas, result.RowsAffected, result.Error
}

func CreateTest(data TbUser) (int64, error) {
	db := orm.Eloquent
	result := db.Create(&data)
	return data.Id, result.Error
}

func FindTest(id int64) (TbUser, error) {
	var model TbUser
	db := orm.Eloquent
	result := db.First(&model, id)
	return model, result.Error
}

func UpdateTest(data TbUser, id int64) (int64, error) {
	var model TbUser
	db := orm.Eloquent
	row := db.First(&model, id)
	if row.Error == nil {
		result := db.Model(&model).Updates(&data)
		return model.Id, result.Error
	}
	return 0, row.Error
}

func DeleteTest(id int64) (int64, error) {
	var model TbUser
	db := orm.Eloquent
	result := db.Delete(&model, id)
	return result.RowsAffected, result.Error
}
