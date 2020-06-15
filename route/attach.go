package route

import (
	"im/app/controller"
	"net/http"
)

func AttachRoutesRegister() {
	//请求前置函数
	http.HandleFunc("/attach/upload", controller.Upload)
	//请求前置函数
	http.HandleFunc("/attach/oss_upload", controller.UploadOss)
}