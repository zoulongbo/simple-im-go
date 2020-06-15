package service

import (
	"errors"
	"fmt"
	"im/app/model"
	"im/app/util"
	"math/rand"
	"time"
)

type UserService struct {

}

//注册函数
func (s *UserService) Register(mobile, plainPwd, nickname, avatar, sex string) (user model.User, err error)  {
	//检测手机号是否存在
	tmp := model.User{}
	_, err = DbEngine.Where("mobile=? ",mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	//如果存在 则提示已注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已注册")
	}
	//插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Password = util.MakePassword(plainPwd, tmp.Salt)
	tmp.Createdat =	time.Now()
	tmp.Token=fmt.Sprintf("%08d", rand.Int31())
	//入库操作
	_, err = DbEngine.InsertOne(&tmp)
	if err != nil {
		return tmp,err
	}

	return tmp, nil
}


//登录函数
func (s *UserService) Login(mobile, plainPwd string) (user model.User, err error)  {
	//检测手机号是否存在
	tmp := model.User{}
	_, err = DbEngine.Where("mobile=? ",mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//不存在提示
	if tmp.Id <= 0 {
		return tmp, errors.New("用户不存在")
	}
	//密码不正确 提示
	if !util.ValidatePassword(plainPwd, tmp.Salt, tmp.Password) {
		return tmp, errors.New("密码不正确")
	}
	//刷新token
	tmp.Token = util.Md5Encode(fmt.Sprintf("%08d", rand.Int31()))

	DbEngine.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}

//用户id获取用户
func (s *UserService) Find(userId int64) (user model.User) {
	//用户是否存在
	tmp := model.User{}
	DbEngine.ID(userId).Get(&tmp)
	return tmp
}