package request

import "time"

type ChangeSegReq struct {
	ToAdd    []string `json:"to_add"`
	ToRemove []string `json:"to_remove"`
	ID       int32    `json:"id"`
}

type UserReq struct {
	ID int32 `json:"id"`
}

type SegmentReq struct {
	Slug string `json:"slug"`
}

type RecordsReq struct {
	ID   int32     `json:"id"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
