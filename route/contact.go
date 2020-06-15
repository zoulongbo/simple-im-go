package route

import (
	"im/app/controller"
	"net/http"
)

func ContactRoutesRegister() {
	http.HandleFunc("/contact/community", controller.LoadCommunity)

	http.HandleFunc("/contact/friend", controller.LoadFriend)

	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)

	http.HandleFunc("/contact/addfriend", controller.AddFriend)
}
