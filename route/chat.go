package route

import (
	"im/app/controller"
	"net/http"
)

func ChatRoutesRegister() {
	//请求前置函数
	http.HandleFunc("/chat", controller.Chat)
}