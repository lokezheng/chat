package model

import "time"

type Message struct {
	Mid        int       `xorm:"pk autoincr int(11)" form:"mid" json:"mid"`
	Id         int64     `xorm:"int(11)" json:"id,omitempty" form:"id"`                       //消息ID
	Userid     int64     `xorm:"int(11)" json:"userid,omitempty" form:"userid"`               //谁发的
	Cmd        int       `xorm:"int(11)" json:"cmd,omitempty" form:"cmd"`                     //群聊还是私聊
	Dstid      int64     `xorm:"int(11)" json:"dstid,omitempty" form:"dstid"`                 //对端用户ID/群ID
	Media      int       `xorm:"int(11)" json:"media,omitempty" form:"media"`                 //消息按照什么样式展示
	Content    string    `xorm:"varchar(1000)" json:"content,omitempty" form:"content"`       //消息的内容
	NewContent string    `xorm:"varchar(1000)" json:"newContent,omitempty" form:"newContent"` //消息的内容
	Pic        string    `xorm:"varchar(255)" json:"pic,omitempty" form:"pic"`                //预览图片
	Url        string    `xorm:"varchar(255)" json:"url,omitempty" form:"url"`                //服务的URL
	Memo       string    `xorm:"varchar(255)" json:"memo,omitempty" form:"memo"`              //简单描述
	Amount     int       `xorm:"int(11)" json:"amount,omitempty" form:"amount"`               //其他和数字相关的
	Createat   time.Time `xorm:"datetime" form:"createat" json:"createat"`                    // 创建时间
}
