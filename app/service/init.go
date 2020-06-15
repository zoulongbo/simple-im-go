package service

import (
	"errors"
	"fmt"
	"im/app/model"
	"im/config"
	"log"
	"xorm.io/xorm"
)

//全局变量 数据库引擎
var DbEngine *xorm.Engine

//main函数运营时 会自动运行init  相当于php construct()
func init() {
	driveName := config.DbDriver
	DsnName := config.DbUsername + ":" + config.DbPassword + "@(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbDatabase + "?charset=" + config.DbCharSet
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driveName, DsnName)
	if err != nil && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//打印sql
	//DbEngine.ShowSQL(true)
	//最大连接数
	DbEngine.SetMaxOpenConns(2)
	//自动user
	DbEngine.Sync2(new(model.User))
	DbEngine.Sync2(new(model.Contact))
	DbEngine.Sync2(new(model.Community))

	fmt.Println("init database ok")
}
