package service

import (
	"chat/model"
	"chat/util"
	"errors"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {
	UserInfo model.User
}

//登录函数
func (s *UserService) Login(
	mobile, //手机
	plainpwd string,
	isAdmin int,
) (user model.User, err error) {
	//首先通过手机号查询用户
	tmp := model.User{}
	_, err = DbEngin.Where("mobile = ?", mobile).Get(&tmp)
	//如果没有找到就注册
	if tmp.Id == 0 {
		//否则拼接插入数据
		tmp.Mobile = mobile
		tmp.Nickname = mobile
		tmp.IsAdmin = isAdmin

		tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
		tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
		tmp.Createat = time.Now()
		tmp.Updateat = time.Now()
		tmp.Loginat = time.Now()
		tmp.Online = 1
		//token 可以是一个随机数
		tmp.Token = fmt.Sprintf("%08d", rand.Int31())
		_, err = DbEngin.InsertOne(&tmp)
		//前端恶意插入特殊字符
		//数据库连接操作失败
		s.UserInfo = tmp
		return tmp, err
	}
	//查询到了比对密码
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	//刷新token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	tmp.Online = 1
	tmp.Loginat = time.Now()
	s.UserInfo = tmp
	//返回数据
	_, err = DbEngin.Where(" id = ?", tmp.Id).Cols("token,online,loginat").Update(&tmp)
	return tmp, err
}

//退出登录
func (s *UserService) LoginOut(
	id int,
	token string,
) (status int, err error) {

	tmp := model.User{}
	_, err = DbEngin.Where("id = ? and token = ?", id, token).Get(&tmp)
	if tmp.Id == 0 {
		return -1, err
	}
	tmp.Token = ""
	tmp.Online = 0
	tmp.Loginoutat = time.Now()
	//返回数据
	_, err = DbEngin.Where(" id = ?", tmp.Id).Cols("token,online,loginat").Update(&tmp)
	return 0, err
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	tmp := model.User{}
	DbEngin.Where("id = ?", userId).Get(&tmp)
	return tmp
}

func (s *UserService) FindByUserName(userName string) (user model.User) {
	tmp := model.User{}
	DbEngin.Where("mobile = ?", userName).Get(&tmp)
	return tmp
}
