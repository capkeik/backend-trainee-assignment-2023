package model

type Segment struct {
	// using id to store in joining table small ints instead of varchars which take up much more memory
	ID    int32  `gorm:"primaryKey;autoIncrement;not null"`
	Slug  string `json:"slug" gorm:"type:varchar(255);not null"`
	Users []User `gorm:"many2many:user_segments;"`
}
