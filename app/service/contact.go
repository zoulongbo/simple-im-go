package service

import (
	"errors"
	"im/app/model"
	"time"
)

func main() {

}

type ContactService struct {
}

//自动添加好友
func (service *ContactService) AddFriend(
	userId, //用户id 10086,
	dstId int64) error {
	//如果加自己
	if userId == dstId {
		return errors.New("不能添加自己为好友啊")
	}
	//判断是否已经加了好友
	tmp := model.Contact{}
	//查询是否已经是好友
	// 条件的链式操作
	DbEngine.Where("owner_id = ?", userId).
		And("dst_id = ?", dstId).
		And("cate = ?", model.CONCAT_CATE_USER).
		Get(&tmp)
	//获得1条记录
	//count()
	//如果存在记录说明已经是好友了不加
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}
	//事务,
	session := DbEngine.NewSession()
	session.Begin()
	//插自己的
	_, e2 := session.InsertOne(model.Contact{
		OwnerId:   userId,
		DstObj:    dstId,
		Cate:      model.CONCAT_CATE_USER,
		CreatedAt: time.Now(),
	})
	//插对方的
	_, e3 := session.InsertOne(model.Contact{
		OwnerId:   dstId,
		DstObj:    userId,
		Cate:      model.CONCAT_CATE_USER,
		CreatedAt: time.Now(),
	})
	//没有错误
	if e2 == nil && e3 == nil {
		//提交
		session.Commit()
		return nil
	} else {
		//回滚
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

func (service *ContactService) SearchComunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	DbEngine.Where("owner_id = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.DstObj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	DbEngine.In("id", comIds).Find(&coms)
	return coms
}

func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	//todo 获取用户全部群ID
	cons := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	DbEngine.Where("owner_id = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&cons)
	for _, v := range cons {
		comIds = append(comIds, v.DstObj)
	}
	return comIds
}

//加群
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		OwnerId: userId,
		DstObj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	DbEngine.Get(&cot)
	if cot.Id == 0 {
		cot.CreatedAt = time.Now()
		_, err := DbEngine.InsertOne(cot)
		return err
	} else {
		return nil
	}

}

//建群
func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.OwnerId == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		OwnerId: comm.OwnerId,
	}
	num, err := DbEngine.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	} else {
		comm.Createdat = time.Now()
		session := DbEngine.NewSession()
		session.Begin()
		_, err = session.InsertOne(&comm)
		if err != nil {
			session.Rollback()
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				OwnerId:   comm.OwnerId,
				DstObj:    comm.Id,
				Cate:      model.CONCAT_CATE_COMUNITY,
				CreatedAt: time.Now(),
			})
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		return com, err
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) []model.User {
	cons := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngine.Where("owner_id = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&cons)
	for _, v := range cons {
		objIds = append(objIds, v.DstObj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngine.In("id", objIds).Find(&coms)
	return coms
}
