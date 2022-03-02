# 郑重声明

# 本系统不得用作生产环境。如因无视本声明，将本系统代码部署到生产环境，导致出现各种损失，本人不承担任何责任！

# 安装方法

本系统升级到golang1.15,请开启如下支持

```
#开启go mod支持
export GO111MODULE=on
#使用代理
export GOPROXY=https://goproxy.io

```

## 1.下载项目

```bash
git clone https://github.com/lokezheng/chat.git
```

## 2.项目配置，非常重要

### 2.1 数据库配置

修改service/init.go 中数据库配置文件

```cgo
const (
	driveName = "mysql"  //数据库类型,不要动
	dsName    = "root:root@(127.0.0.1:3306)/chat?charset=utf8"  //tech-chat是数据库名称,请先创建
	showSQL   = true  //是否显示sql语句
	maxCon    = 10  //最大连接数
	NONERROR  = "noerror" //一个字符串标记常量
)
```

为你自己的数据库以及密码,格式如下

```
用户名:密码@(ip:port)/数据库名称?charset=utf8
```

### 2.2 配置子网掩码,防火墙开放3000

修改ctrl/chat.go 175行左右

```cgo
func udpsendproc() {
	log.Println("start udpsendproc")
	//todo 使用udp协议拨号
	con, err := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 0, 255),
			Port: 3000,
	})
    //....
}

```

其中`IP:net.IPv4(192, 168, 0, 255)`, 改为你当前应用所在服务器的子网掩码, 举个简单一点的例子,比如当前应用所安装环境是`192.168.3.1`
，则需要修改参数为`net.IPv4(192, 168, 3, 255)`
`Port: 3000`为通信端口。本系统依赖于UPD进行分布式部署。因此需要在防火墙内开放该端口。

### 2.3 分布式部署

本系统支持分布式部署,要求是将当前应用部署在同一个网段中。代码修改同2.3

### 2.4 页面入口地址

```
http://127.0.0.1:8899/user/login.shtml
```

### 2.5 中方分词库gojieba

在win系统依赖 gcc库 会出现编译问题 建议linux跑

## 3.依赖包安装

使用go mod 自动处理安装包

## 4. 操作说明

登录选择GM，可以使用GM命令

1. /stats [username]
2. /popular n (n为房间Id)

## 5，引用第三方库

	github.com/aliyun/aliyun-oss-go-sdk 
	github.com/baiyubin/aliyun-sts-go-sdk      阿里云OSS
	github.com/go-sql-driver/mysql        
	github.com/go-xorm/xorm                     xorm
	github.com/gorilla/websocket                websocket

	github.com/satori/go.uuid                   uuid
	github.com/tchap/go-patricia                基数树用与脏词过滤
	github.com/yanyiwu/gojieba                   中文分词

