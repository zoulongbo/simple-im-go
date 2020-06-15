package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"im/route"
	"log"
	"net/http"
)


//模版加载
func registerTemplate() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplName, nil)
		})

	}

	//单个模版渲染
	/*http.HandleFunc("/user/login.shtml", func(writer http.ResponseWriter, request *http.Request) {
		//解析
		tpl, err := template.ParseFiles("view/user/login.html")
		if err != nil {
			log.Fatal(err.Error())		//打印并直接退出
		}
		tpl.ExecuteTemplate(writer, "/user/login.shtml", nil)
	})*/
}

func main() {
	//路由注册
	route.RoutesRegister()
	//静态文件加载
	http.Handle("/public/", http.FileServer(http.Dir(".")))
	//模版加载
	registerTemplate()
	//启动http服务
	http.ListenAndServe(":8080", nil)
}
