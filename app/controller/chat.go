package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"im/app/args"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

//客户端与node映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwLocker sync.RWMutex

// ws://127.0.0.1/chat?id=1&token=xxxxx
func Chat(write http.ResponseWriter, request *http.Request) {
	//参数获取校验 此种方式获取到的参数为字符串
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isValidDa := checkToken(userId, token)
	if !isValidDa {
		return
	}
	//获取ws链接
	wsConn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValidDa
		},
	}).Upgrade(write, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//userId 和node绑定关系
	node := &Node{
		Conn:      wsConn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//群聊关系绑定
	comIds := contactService.SearchComunityIds(userId)
	for _,v := range comIds {
		node.GroupSets.Add(v)
	}

	//用户连接关系绑定
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	go sendProc(node)

	go recvProc(node)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	sendMessage(userId, []byte("hello world"))

}

func checkToken(userId int64, token string) bool {
	user := userService.Find(userId)
	return user.Token == token
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}

		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		// 处理data
		dispatch(data)
		fmt.Printf("recv<=%s", data)
	}
}

//发送消息
func sendMessage(userId int64, msg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[userId]
	rwLocker.Unlock()
	if ok {
		node.DataQueue <- msg
	}
}

func dispatch(data []byte) {
	msg := args.Message{}
	//数据解析成json
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//根据cmd处理逻辑
	switch msg.Cmd {
	case args.CMD_SINGLE_MSG:
		sendMessage(msg.DstId, data)
	case args.CMD_ROOM_MSG:
		//TODO 群聊
		for _, v:= range clientMap {
			if v.GroupSets.Has(msg.DstId) {
				v.DataQueue <- data
			}
		}
	case args.CMD_HEART:
		//啥都不做

	}

}
