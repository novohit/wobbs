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

// GetUserInfoByID 用户详情
// @Summary 用户详情
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /user/info [get]
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

// Register 用户注册
// @Summary 用户注册
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body dto.RegisterDTO true "JSON数据"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /user/register [post]
func Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := ctx.ShouldBind(&registerDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	logic.Register(registerDTO)

	common.Success(ctx, nil)
}

// Login 用户登录
// @Summary 用户登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param data body dto.LoginDTO true "JSON数据"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /user/login [post]
func Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	tokens := logic.Login(loginDTO, ctx)
	common.Success(ctx, tokens)
}
