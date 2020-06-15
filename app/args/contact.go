package args

type ContactArg struct {
	PageArg
	UserId int64 	`json:"user_id" form:"user_id"`
	DstId int64 	`json:"dst_id" form:"dst_id"`
}
