package model

import "time"

type Contact struct {
	Id 			int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//谁的10000
	OwnerId 	int64 `xorm:"bigint(20)" form:"owner_id" json:"owner_id"` // 记录是谁的
	//对端,10001
	DstObj		int64 `xorm:"bigint(20)" form:"dst_obj" json:"dst_obj"` // 对端信息
	//
	Cate 		int    `xorm:"int(11)" form:"cate" json:"cate"`      // 什么类型
	Memo 		string `xorm:"varchar(120)" form:"memo" json:"memo"` // 备注
	//
	CreatedAt 	time.Time `xorm:"datetime" form:"created_at" json:"created_at"` // 创建时间
}

const (
	CONCAT_CATE_USER = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)
