package request

import (
	"time"
)

type ChangeSegReq struct {
	ToAdd    []string `json:"to_add" validate:"required"`
	ToRemove []string `json:"to_remove" validate:"required"`
	ID       int32    `json:"id" validate:"required"`
}

type UserReq struct {
	ID int32 `json:"id" validate:"required"`
}

type SegmentReq struct {
	Slug string `json:"slug" validate:"required"`
}

type RecordsReq struct {
	ID   int32     `json:"id" validate:"required"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
