package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type M struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, msg, nil)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, msg, data)
}

func RespOkList(w http.ResponseWriter, lists interface{}, total interface{}) {
	//分页数目,
	RespList(w, 0, lists, total)
}

// 数据返回格式统一方法
func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	m := M{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//把结构体转化为json
	ret, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
	}
	//默认返回text/html
	w.Header().Set("Content-Type", "application/json")
	//默认返回200
	w.WriteHeader(http.StatusOK)

	w.Write(ret)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	m := M{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//输出
	//定义一个结构体
	//满足某一条件的全部记录数目
	//测试 100
	//20
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	w.Write(ret)
}
