package model

type UserSegment struct {
	UserID    int32 `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	SegmentID int32 `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
}
