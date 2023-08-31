package model

type User struct {
	// int32 because capacity of postgresql's int and serial is 4 bytes
	ID       int32     `gorm:"primaryKey;autoIncrement:false"`
	Segments []Segment `gorm:"many2many:user_segments;"`
	Records  []Record  `gorm:"foreignKey:UserID;references:ID;onDelete:CASCADE"`
}
