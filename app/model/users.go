package model

import "time"

func CreateUsersFactory() *UserModel {
	return &UserModel{BaseModel: BaseModel{DB: UseDbConn("")}}
}

type UserModel struct {
	BaseModel
	Openid  string    `gorm:"column:openid;type:varchar(32);not null;default:''" json:"openid"`    // 微信openid
	Mobile  string    `gorm:"column:mobile;type:char(11);not null;default:''" json:"mobile"`       // 手机号
	Name    string    `gorm:"column:name;type:varchar(64);not null;default:''" json:"name"`        // 昵称
	City    string    `gorm:"column:city;type:varchar(64);not null;default:''" json:"city"`        // 城市
	Sex     bool      `gorm:"column:sex;type:tinyint(1);not null;default:0" json:"sex"`            // sex1 男 2女 0未知
	LoginAt time.Time `gorm:"column:login_at;type:TIMESTAMP;not null" json:"loginAt"`              // 登录时间
	LoginIP string    `gorm:"column:login_ip;type:varchar(32);not null;default:''" json:"loginIp"` // 登录IP
	Status  bool      `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`      // 状态1正常 0禁用
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) List(page int, size int) (items []UserModel) {
	u.Model(u).Scopes(Paginate(page, size)).Debug().Find(&items)
	return items
}
