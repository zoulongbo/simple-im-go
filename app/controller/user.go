package controller

import (
	"fmt"
	"im/app/model"
	"im/app/service"
	"im/app/util"
	"log"
	"math/rand"
	"net/http"
)

var userService service.UserService

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	//io.WriteString(writer, "hello world")
	//业务校验
	//数据库
	//逻辑处理
	//restapi 返回
	request.ParseForm()
	mobile := request.Form.Get("mobile")
	password := request.Form.Get("password")
	user, err := userService.Login(mobile, password)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		data := make(map[string]interface{})
		data["id"] = user.Id
		data["token"] = user.Token
		data["nickname"] = user.Nickname
		data["avatar"] = user.Avatar
		util.RespOk(writer, data, "")
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	//restapi 返回
	request.ParseForm()
	mobile := request.Form.Get("mobile")
	password := request.Form.Get("password")
	nickname := fmt.Sprintf("user%06d", rand.Int31n(100000))
	avatar := ""
	sex := model.SEX_UNKNOW
	user, err := userService.Register(mobile, password, nickname, avatar, sex)
	log.Println(err)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}
