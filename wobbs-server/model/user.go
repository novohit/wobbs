package model

// User 前端接收int64数字失真解决方案 打上Tag json:"user_id,string" 序列化时自动将int64转换为string
// 同时反序列化时会自动将string转换为int64赋值为struct
type User struct {
	UserID   int64  `gorm:"index:idx_user_id;unique;not null" json:"user_id,string"`
	Username string `gorm:"index:idx_username;unique;type:varchar(256);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"-"`
	Email    string `gorm:"type:varchar(256)" json:"email"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(8) comment 'male:男 female:女'" json:"gender"`
	Age      int    `gorm:"column:age;type:int comment 'male:男 female:女'" json:"age"`
	BaseModel
}
