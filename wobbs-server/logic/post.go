package logic

import (
	"wobbs-server/config"
	"wobbs-server/dto"
	"wobbs-server/model"
)

func GetPostDetail(pid int32) model.Post {
	db := config.GetDB()
	var post model.Post
	db.Where(pid).Find(&post)
	return post
}

func CreatePost(userId int64, post dto.PostDTO) {
	db := config.GetDB()
	db.Create(&model.Post{AuthorID: userId,
		CategoryID: post.CategoryID,
		Title:      post.Title,
		Content:    post.Content,
		Status:     post.Status,
	})
}
