package api

import (
	"chat/model"
	"chat/service"
	"chat/util"
	"net/http"
	"strconv"
)

func UserLogin(writer http.ResponseWriter,
	request *http.Request) {
	//数据库操作
	//逻辑处理
	//restapi json/xml返回
	//1.获取前端传递的参数
	//mobile,passwd
	//解析参数
	//如何获得参数
	//解析参数
	request.ParseForm()

	mobile := request.PostForm.Get("mobile")
	isAdmin := request.PostForm.Get("isAdmin")
	passwd := request.PostForm.Get("passwd")

	i, err := strconv.Atoi(isAdmin)
	if err != nil {
		util.RespFail(writer, err.Error())
	}
	//模拟
	user, err := userService.Login(mobile, passwd, i)

	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}

}

var userService service.UserService

//解析一下
func FindUserById(writer http.ResponseWriter,
	request *http.Request) {
	var user model.User
	util.Bind(request, &user)
	user = userService.Find(user.Id)
	if user.Id == 0 {
		util.RespFail(writer, "该用户不存在")
	} else {
		util.RespOk(writer, user, "")

	}

}
