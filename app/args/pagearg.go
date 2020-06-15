package args

import (
	"fmt"
	"time"
)

type PageArg struct {
	//从哪页开始
	PageFrom int `json:"page_from" form:"page_from"`
	//每页大小
	PageSize int `json:"page_size" form:"page_size"`
	//关键词
	Keyword string `json:"keyword" form:"keyword"`
	//asc：“id”  id asc
	Asc  string `json:"asc" form:"asc"`
	Desc string `json:"desc" form:"desc"`
	//
	Name string `json:"name" form:"name"`
	//
	UserId int64 `json:"user_id" form:"user_id"`
	//dst_id
	DstId int64 `json:"dst_id" form:"dst_id"`
	//时间点1
	DateFrom time.Time `json:"data_from" form:"data_from"`
	//时间点2
	DateTo time.Time `json:"date_to" form:"date_to"`
	//
	Total int64 `json:"total" form:"total"`
}

func (p *PageArg) GetPageSize() int {
	if p.PageSize == 0 {
		return 100
	} else {
		return p.PageSize
	}

}
func (p *PageArg) GetPageFrom() int {
	if p.PageFrom < 0 {
		return 0
	} else {
		return p.PageFrom
	}
}

func (p *PageArg) GetOrderBy() string {
	if len(p.Asc) > 0 {
		return fmt.Sprintf(" %s asc", p.Asc)
	} else if len(p.Desc) > 0 {
		return fmt.Sprintf(" %s desc", p.Desc)
	} else {
		return ""
	}
}
