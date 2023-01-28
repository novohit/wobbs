package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wobbs-server/common"
	"wobbs-server/config"
	"wobbs-server/dto"
	"wobbs-server/logic"
)

func GetUserInfoByID(ctx *gin.Context) {
	/*uid := util.StringToInt(ctx.Query("uid"))
	res := service.GetUserInfoByIDService(uid)
	response.HandleResponse(ctx, res)*/
	value, _ := ctx.Get("userId")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": value,
	})
}

func Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := ctx.ShouldBind(&registerDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	logic.Register(registerDTO)

	common.Success(ctx, nil)
}

func Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := ctx.ShouldBind(&loginDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	token := logic.Login(loginDTO)
	common.Success(ctx, token)
}
