package model

import "time"

type User struct {
	//用户ID
	Id       int64  `xorm:"pk autoincr int(11)" form:"id" json:"id"`
	Mobile   string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `xorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 昵称
	IsAdmin  int    `xorm:"int(1)" form:"is_admin" json:"is_admin"`      //是否是gm

	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`   //加盐随机字符串6
	Online int    `xorm:"int(1)" form:"online" json:"online"` //是否在线

	Token      string    `xorm:"varchar(40)" form:"token" json:"token"`        //前端鉴权因子,本来应该使用jwt之类的
	Memo       string    `xorm:"varchar(140)" form:"memo" json:"memo"`         // 备注信息
	Createat   time.Time `xorm:"datetime" form:"createat" json:"createat"`     // 创建时间
	Updateat   time.Time `xorm:"datetime" form:"updateat" json:"updateat"`     // 更新时间
	Loginat    time.Time `xorm:"datetime" form:"loginat" json:"loginat"`       // 最后登录时间
	Loginoutat time.Time `xorm:"datetime" form:"loginoutat" json:"loginoutat"` // 退出登录时间
}
