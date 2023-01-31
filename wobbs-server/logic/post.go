package logic

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
	"wobbs-server/common"
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

func GetPostList(page int, pageSize int) []model.Post {
	db := config.GetDB()
	posts := make([]model.Post, 0)
	db.Preload("User").Preload("Category").Scopes(config.Paginate(page, pageSize)).Find(&posts)
	return posts
}

func CreatePost(userId int64, dto dto.PostDTO) {
	db := config.GetDB()
	pipeline := config.RDB.TxPipeline()
	ctx := context.Background()
	post := model.Post{
		AuthorID:   userId,
		CategoryID: dto.CategoryID,
		Title:      dto.Title,
		Content:    dto.Content,
		Status:     dto.Status,
	}
	// 入库
	db.Create(&post)
	// redis记录帖子发布时间
	pipeline.ZAdd(ctx, common.KeyPostTimeZSet, &redis.Z{
		Score:  float64(post.CreateTime.Unix()),
		Member: post.ID,
	})
	// 设置帖子初始分值为帖子发布时间
	pipeline.ZAdd(ctx, common.KeyPostScoreZSet, &redis.Z{
		Score:  float64(post.CreateTime.Unix()),
		Member: post.ID,
	})
	if _, err := pipeline.Exec(ctx); err != nil {
		zap.L().Error("创建帖子 redis操作异常")
		return
	}
	zap.L().Info("redis记录发帖时间 " + fmt.Sprintf("%d", post.CreateTime.Unix()))
}

const (
	OneWeekTime = 7 * 24 * 3600
	Score       = 432
)

// PostVoting 一天86400s 定义 200张投票=一天时间的分数
// 分数计算：帖子创建时间+点赞数*200
// TODO 如何做持久化？
func PostVoting(userId int64, vote dto.VoteDTO) {
	userIdStr := strconv.FormatInt(userId, 10)
	postIdStr := strconv.FormatInt(int64(vote.PostID), 10)

	rdb := config.RDB
	ctx := context.Background()
	// 获取当前帖子的发表时间
	//postTime := rdb.ZScore(ctx, common.KeyPostTimeZSet, postIdStr).Val()
	//if float64(time.Now().Unix())-postTime > OneWeekTime {
	//	panic(common.NewCustomError(common.CodeVoteTimeExpired))
	//}
	// 1. 查询当前用户是否有投票记录

	t := rdb.HGet(ctx, common.KeyPostVotedPrefix+postIdStr, userIdStr).Val()
	record, _ := strconv.Atoi(t)
	// 2. 计算差值
	diff := vote.Type - record
	// 3. 更新当前帖子的分值
	rdb.ZIncrBy(ctx, common.KeyPostScoreZSet, float64(diff*Score), postIdStr)
	// 4. 更新用户的投票记录
	rdb.HSet(ctx, common.KeyPostVotedPrefix+postIdStr, userIdStr, strconv.Itoa(vote.Type))
}
