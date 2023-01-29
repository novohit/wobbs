package api

import (
	"github.com/gin-gonic/gin"

	"wobbs-server/common"
	"wobbs-server/logic"
)

func GetCommunityCategory(ctx *gin.Context) {
	category := logic.GetCommunityCategory()
	common.Success(ctx, category)
}
