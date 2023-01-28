package logic

import (
	"errors"
	"wobbs-server/common"
	"wobbs-server/config"
	"wobbs-server/dto"
	"wobbs-server/model"
	"wobbs-server/pkg/jwt"
	"wobbs-server/pkg/snowflake"
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

func Login(user dto.LoginDTO) string {
	dbUser := FindUserByUsername(user.Username)
	if dbUser.ID == 0 {
		panic(common.NewCustomError(common.CodeInvalidPassword))
	}
	if dbUser.Password != user.Password {
		panic(common.NewCustomError(common.CodeInvalidPassword))
	}
	token, err := jwt.AccessToken(dbUser.UserID)
	if err != nil {
		panic(err)
	}
	return token
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
