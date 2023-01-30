package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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
	userId, _ := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": logic.GetUserById(userId),
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
	tokens := logic.Login(loginDTO, ctx)
	common.Success(ctx, tokens)
}
