package request

type ChangeSegReq struct {
	ToAdd    []string `json:"to_add"`
	ToRemove []string `json:"to_remove"`
	Id       int32    `json:"id"`
}

type UserReq struct {
	ID int32 `json:"id"`
}
