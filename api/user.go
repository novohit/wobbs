package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfoByID(ctx *gin.Context) {
	/*uid := util.StringToInt(ctx.Query("uid"))
	res := service.GetUserInfoByIDService(uid)
	response.HandleResponse(ctx, res)*/
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}
