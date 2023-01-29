package logic

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wobbs-server/common"
	"wobbs-server/config"
	"wobbs-server/dto"
	"wobbs-server/model"
	"wobbs-server/pkg/jwt"
	"wobbs-server/pkg/snowflake"
	"wobbs-server/vo"
)

func Register(user dto.RegisterDTO) {

	DB := config.GetDB()

	if isUserExist(user.Username) {
		panic(common.NewCustomError(common.CodeUserExist))
	}

	newUser := model.User{
		UserID:   snowflake.GenerateID(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}
	result := DB.Create(&newUser)
	if result.RowsAffected == 0 {
		panic(errors.New("创建失败"))
	}
}

func Login(user dto.LoginDTO, ctx *gin.Context) vo.Tokens {
	dbUser := FindUserByUsername(user.Username)
	if dbUser.ID == 0 {
		panic(common.NewCustomError(common.CodeInvalidPassword))
	}
	if dbUser.Password != user.Password {
		panic(common.NewCustomError(common.CodeInvalidPassword))
	}
	accessToken, err := jwt.AccessToken(dbUser.UserID)
	if err != nil {
		panic(err)
	}
	refreshToken, err := jwt.RefreshToken(dbUser.UserID)
	if err != nil {
		panic(err)
	}
	// 将access_token 存入redis中 限制同一用户同一IP 同一时间只能登录一个设备
	// key user:token:user_id:IP value access_token
	config.RDB.Set(context.Background(), common.ConstUserTokenPrefix+strconv.FormatInt(dbUser.UserID, 10)+":"+ctx.RemoteIP(), accessToken, 2*time.Hour)
	return vo.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func isUserExist(username string) bool {
	DB := config.GetDB()
	var user model.User
	result := DB.Where(&model.User{Username: username}).Find(&user)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func FindUserByUsername(username string) model.User {
	DB := config.GetDB()
	var user model.User
	DB.Where(&model.User{Username: username}).Find(&user)
	return user
}
