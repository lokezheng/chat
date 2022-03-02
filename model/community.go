package model

import "time"

type Community struct {
	Id int64 `xorm:"pk autoincr int(11)" form:"id" json:"id"`
	//名称
	Name string `xorm:"varchar(30)" form:"name" json:"name"`
	//群主ID
	Ownerid int64 `xorm:"bigint(11)" form:"ownerid" json:"ownerid"`
	//群logo
	Icon string `xorm:"varchar(250)" form:"icon" json:"icon"`
	//como
	Cate int `xorm:"int(11)" form:"cate" json:"cate"`
	//描述
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"`
	//创建时间
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
	//更新时间
	Updateat time.Time `xorm:"datetime" form:"updateat" json:"updateat"`
}

const (
	COMMUNITY_CATE_COM = 0x01
)
