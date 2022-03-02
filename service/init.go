package service

import (
	"chat/model"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

const (
	driveName     = "mysql"
	dsName        = "root:123456@(127.0.0.1:3306)/chat?charset=utf8"
	showSQL       = true
	maxCon        = 10
	NONERROR      = "noerror" //没有错误
	BlacklistFile = "mnt/globalBlacklist.golden"
)

var DbEngin *xorm.Engine

//初始化数据库
func init() {
	getDB()
}

func getDB() {
	err := errors.New(NONERROR)
	DbEngin, err = xorm.NewEngine(driveName, dsName)
	if nil != err && NONERROR != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(showSQL)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(maxCon)

	//自动User
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community),
		new(model.Message))
	fmt.Println("init data base ok")
}
