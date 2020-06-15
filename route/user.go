package route

import (
	"im/app/controller"
	"net/http"
)

func UserRoutesRegister() {
	//请求前置函数
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
}
