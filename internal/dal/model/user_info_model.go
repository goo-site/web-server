// Code generated by ECode v1.3.1. DO NOT EDIT.
package model

//go:generate msgp --tests=false -o user_info_msgp.go
type UserInfo struct {
	ID         *int64 `gorm:"column:id;primaryKey" json:"id"`                        // id
	UserID     string `gorm:"column:user_id;uniqueIndex:user_id_key" json:"user_id"` // 用户id
	Password   string `gorm:"column:password" json:"password"`                       // 用户密码
	Email      string `gorm:"column:email" json:"email"`                             // 用户邮箱
	Nickname   string `gorm:"column:nickname" json:"nickname"`                       // 用户昵称
	Permission int8   `gorm:"column:permission" json:"permission"`                   // 用户权限：1、root权限 2、普通权限
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`                 // 创建时间
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`                 // 更新时间
}

func (UserInfo) TableName() string {
	return "user_info"
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}