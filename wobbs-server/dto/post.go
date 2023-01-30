package dto

type PostDTO struct {
	CategoryID int32  `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Status     int    `json:"status" binding:"required"`
}
