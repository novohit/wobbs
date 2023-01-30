package model

type Post struct {
	AuthorID   int64    `gorm:"index:idx_author_id;column:author_id;type:bigint(20);not null" json:"author_id"`
	CategoryID int32    `gorm:"column:category_id;not null" json:"category_id"`
	Title      string   `gorm:"column:title;type:varchar(512)" json:"title"`
	Content    string   `gorm:"column:content" json:"content"`
	Status     int      `gorm:"column:status;type:tinyint;comment:文章状态" json:"status"`
	User       User     `gorm:"foreignKey:UserID;references:AuthorID" json:"user"`
	Category   Category `gorm:"foreignKey:ID;references:CategoryID;" json:"category"`
	BaseModel
}
