package entity

// 用户信息表
type UserInfo struct {
	Id         uint64 `gorm:"column:id" json:"id"`                   //用户id
	Account    string `gorm:"column:account" json:"account"`         //用户账号
	Password   string `gorm:"column:password" json:"password"`       //用户密码
	Nickname   string `gorm:"column:nickname" json:"nickname"`       //用户昵称
	Permission int    `gorm:"column:permission" json:"permission"`   //用户权限：1、root权限 2、普通权限
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` //创建时间
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"` //更新时间
}

func (UserInfo) TableName() string {
	return "user_info"
}
