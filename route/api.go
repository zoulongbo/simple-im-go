package route

//路由注册
func RoutesRegister() {
	//用户相关路由注册
	UserRoutesRegister()
	//联系人相关路由注册
	ContactRoutesRegister()
	//聊天路由注册
	ChatRoutesRegister()
	//附件操作路由注册
	AttachRoutesRegister()
}