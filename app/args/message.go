package args


type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"` 					//消息ID
	UserId  int64  `json:"user_id,omitempty" form:"user_id"` 		//谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`		 		//群聊还是私聊
	DstId   int64  `json:"dst_id,omitempty" form:"dst_id"`			//对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"` 			//消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` 		//消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"` 				//预览图片
	Url     string `json:"url,omitempty" form:"url"` 				//服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"` 				//简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"` 			//其他和数字相关的
}

const (
	CMD_SINGLE_MSG = 10						//CMD 单聊
	CMD_ROOM_MSG   = 11						//CMD 群聊
	CMD_HEART      = 0						//CMD 心跳检测

	MEDIA_TYPE_TEXT = 1						//MEDIA 文本
	MEDIA_TYPE_NEWS = 2						//MEDIA 图文
	MEDIA_TYPE_VOICE = 3					//MEDIA 音频
	MEDIA_TYPE_IMG = 4						//MEDIA 图片
	MEDIA_TYPE_RED_PACKAGE = 5				//MEDIA 红包
	MEDIA_TYPE_EMOJ = 6						//MEDIA 表情
	MEDIA_TYPE_LINK = 7						//MEDIA 超链接
	MEDIA_TYPE_VIDEO = 8 					//MEDIA 视频
	MEDIA_TYPE_UDEF	= 100					//自定义

)