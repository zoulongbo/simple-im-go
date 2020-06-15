# simple-im-go

#数据库增删改查一般套路
 ##一、安装初始化 xorm.NewSession(driverName,dataSourceName)
 ##二、定义实体 模型层model或者实体层entity
  ###1、定义与表结构对应对象User

    type User struct {
        Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
        Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
        Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 什么角色
        Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
        Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`   // 什么角色
        Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // 什么角色
        Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`   // 什么角色
        Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
        Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   // 什么角色
        Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`   // 什么角色
        Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   // 什么角色
    }
##三、定义和业务相关的服务 服务层service,专门用来存放数据库业务服务的,如 注册、登录
 ###2、查询单个用户Find,参数userId
     DbEngin.ID(userId).Get(&User)
###3、查询满足某一类条件的Search

    //
    result :=make([]User,0)
    DbEngin.where("mobile=? ",moile).Find(&result)
    DbEngin.where("mobile=? ",moile).Get(&User)
   
###4、创建一条记录Create

    DBengin.InsertOne(&User)
###5、修改某条记录Update

     DBengin.ID(userId).Update(&User)
     // update ... where id = xx
     DBengin.Where("a=? and b=?",a,b).Update(&User)
     DBengin.Where("a=? and b=?",a,b).Cols("nick_name").Update(&User)
###6、删除某条记录Delete

     DBengin.ID(userId).Delete(&User)
###7、MD5加密函数

    import (
        "crypto/md5"
        "encoding/hex"
        "strings"
    )
    
    func Md5Encode(data string) string{
        h := md5.New()
        h.Write([]byte(data)) // 需要加密的字符串为 123456
        cipherStr := h.Sum(nil)
    
        return  hex.EncodeToString(cipherStr)
    
    }
    func MD5Encode(data string) string{
        return strings.ToUpper(Md5Encode(data))
    }
    
    func ValidatePasswd(plainpwd,salt,passwd string) bool{
        return Md5Encode(plainpwd+salt)==passwd
    }
    func MakePasswd(plainpwd,salt string) string{
        return Md5Encode(plainpwd+salt)
    }
##四、控制器层调用

    var userServer server.UserServer
    type UserCtrl struct{}
    
    func (ctrl *UserCtrl)Register(w){
        user = userServer.Register(mobile,passwd)
    }
## 五、消息体
前端websocket发送json字符串

```
    数据：json
    example: 
        消息发送结构体
        1、MEDIA_TYPE_TEXT  1
        {id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
        2、MEDIA_TYPE_News  2
        {id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}
        3、MEDIA_TYPE_VOICE，amount单位秒  3
        {id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}
        4、MEDIA_TYPE_IMG   4 
        {id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}
        5、MEDIA_TYPE_REDPACKAGR //红包amount 单位分  5 
        {id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
        6、MEDIA_TYPE_EMOJ 6
        {id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}
        7、MEDIA_TYPE_Link 7
        {id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}
        8、MEDIA_TYPE_VIDEO 8
        {id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}
        9、MEDIA_TYPE_CONTACT 9
        {id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}
    字段解释：
            1、  userid : 发送者id
            2、  dstid  : 发送对象id
            3、  cmd:     消息用途
            4、  media:   展示方式
            5、  消息内容： content,pic,url
    message结构体:
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
```