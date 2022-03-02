package main

import (
	"net/http"

	v1 "chat/api/v1"
	"fmt"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//注册首页自动跳转
func RegisterIndex() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/user/login.shtml", http.StatusFound) //跳转到百度
	})
}

//注册模板
func RegisterView() {
	//一次解析出全部模板
	tpl, err := template.ParseGlob("view/**/*")
	if nil != err {
		log.Fatal(err)
	}
	//通过for循环做好映射
	for _, v := range tpl.Templates() {
		//
		tplname := v.Name()
		fmt.Println("HandleFunc     " + v.Name())
		http.HandleFunc(tplname, func(w http.ResponseWriter,
			request *http.Request) {
			//
			fmt.Println("parse     " + v.Name() + "==" + tplname)
			err := tpl.ExecuteTemplate(w, tplname, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		})
	}

}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", v1.UserLogin)
	http.HandleFunc("/user/find", v1.FindUserById)
	http.HandleFunc("/contact/loadcommunity", v1.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", v1.LoadFriend)
	http.HandleFunc("/contact/joincommunity", v1.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", v1.CreateCommunity)
	http.HandleFunc("/contact/addfriend", v1.Addfriend)
	http.HandleFunc("/chat", v1.Chat)
	http.HandleFunc("/attach/upload", v1.Upload)

	// 指定目录的静态文件
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))

	RegisterView()
	RegisterIndex()

	fmt.Println("run at :8899")
	http.ListenAndServe(":8899", nil)
}
