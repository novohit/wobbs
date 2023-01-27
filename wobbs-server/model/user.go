package model

type User struct {
	UserID   int64  `gorm:"index:idx_phone;unique;type:bigint(20);not null"`
	Username string `gorm:"index:idx_phone;unique;type:varchar(256);not null"`
	Password string `gorm:"type:varchar(256);not null"`
	Email    string `gorm:"type:varchar(256)"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(8) comment 'male:男 female:女'"`
	Age      int    `gorm:"column:age;type:int comment 'male:男 female:女'"`
	BaseModel
}
