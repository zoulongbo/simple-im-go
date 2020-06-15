package model

import "time"

const  (
	SEX_WOMEN="W"
	SEX_MAN="M"
	SEX_UNKNOW="U"

)

type User struct {
	Id        int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"` //用户id
	Mobile    string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`    //用户手机号
	Password  string    `xorm:"varchar(40)" form:"password" json:"-"`       //用户密码
	Avatar    string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`	//头像
	Sex       string    `xorm:"varchar(2)" form:"sex" json:"sex"`             //性别
	Nickname  string    `xorm:"varchar(20)" form:"nickname" json:"nickname"`  // 昵称
	Salt      string    `xorm:"varchar(10)" form:"salt" json:"-"`             // 盐
	Online    int       `xorm:"int(10)" form:"online" json:"online"`          //是否在线
	Token     string    `xorm:"varchar(40)" form:"token" json:"token"`        // token
	Memo      string    `xorm:"varchar(140)" form:"memo" json:"memo"`         // 什么角色
	Createdat time.Time `xorm:"datetime" form:"created_at" json:"created_at"` // 什么角色
}
