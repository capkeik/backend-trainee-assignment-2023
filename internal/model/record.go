package model

import "time"

type Action string

const (
	ActionAdd    Action = "add"
	ActionRemove Action = "remove"
)

type Record struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	UserID          int32     `gorm:"index;"`
	Slug            string    `gorm:"not null"`
	Action          Action    `gorm:"type:action_enum;not null"`
	ActionTimestamp time.Time `gorm:"default:now();"`
}
