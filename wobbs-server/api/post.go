package api

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"wobbs-server/vo"

	"github.com/gin-gonic/gin"

	"wobbs-server/common"
	"wobbs-server/config"
	"wobbs-server/dto"
	"wobbs-server/logic"
)

func GetPostDetail(ctx *gin.Context) {
	pidStr := ctx.Param("pid")
	if pidStr == "" {
		zap.L().Error("pid")
		common.FailByMsg(ctx, "pid为空")
		return
	}
	pid, err := strconv.ParseInt(pidStr, 10, 32)
	if err != nil {

	}
	detail := logic.GetPostDetail(int32(pid))
	category := logic.GetCategoryById(detail.CategoryID)
	author := logic.GetUserById(detail.AuthorID)
	common.Success(ctx,
		vo.PostDetail{AuthorName: author.Username,
			CategoryName: category.Name,
			Post:         detail})
}

func CreatePost(ctx *gin.Context) {
	var postDTO dto.PostDTO
	if err := ctx.ShouldBind(&postDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	userId, exists := ctx.Get("userId")
	if !exists {
		fmt.Println("未登录")
	}
	logic.CreatePost(userId.(int64), postDTO)
	common.Success(ctx, nil)
}
