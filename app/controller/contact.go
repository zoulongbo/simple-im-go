package controller

import (
	"im/app/args"
	"im/app/service"
	"im/app/util"
	"net/http"
)

var contactService service.ContactService

func LoadCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(request, &arg)
	communitys := contactService.SearchComunity(arg.UserId)
	util.RespOkList(writer, communitys, len(communitys))
}

func LoadFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	err := util.Bind(request, &arg)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	friendList := contactService.SearchFriend(arg.UserId)
	util.RespOkList(writer, friendList, len(friendList))
	return
}

func JoinCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(request,&arg)
	err := contactService.JoinCommunity(arg.UserId,arg.DstId);
	//todo 刷新用户的群组信息
	addGroupId(arg.UserId,arg.DstId)
	if err!=nil{
		util.RespFail(writer,err.Error())
	}else {
		util.RespOk(writer,nil,"")
	}
}

func AddFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	util.Bind(request,&arg)
	//调用service
	err := contactService.AddFriend(arg.UserId,arg.DstId)
	//
	if err!=nil{
		util.RespFail(writer,err.Error())
	}else{
		util.RespOk(writer,nil,"好友添加成功")
	}
}

func addGroupId(userId , gid int64) {
	rwLocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	rwLocker.Unlock()
}
