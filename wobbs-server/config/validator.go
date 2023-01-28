package config

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func InitTranslator(language string) error {
	// 修改gin框架中的validator引擎属性
	validate = binding.Validator.Engine().(*validator.Validate)

	// 注册一个获取json tag的自定义方法
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	zh := zh.New()
	en := en.New()
	// 第一个参数是fallback locale 后面的参数是locales it should support
	uni = ut.New(en, zh, en)
	var ok bool
	trans, ok = uni.GetTranslator(language)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", language)
	}
	switch language {
	case "en":
		en_translations.RegisterDefaultTranslations(validate, trans)
	case "zh":
		zh_translations.RegisterDefaultTranslations(validate, trans)
	default:
		zh_translations.RegisterDefaultTranslations(validate, trans)
	}
	return nil
}

// RegisterDTO.name 删除错误信息前面的RegisterDTO.
func removeTopStruct(errMsg map[string]string) map[string]string {
	res := make(map[string]string)
	for field, msg := range errMsg {
		res[field[strings.Index(field, ".")+1:]] = msg
	}
	return res
}

func ValidateError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	// 如果不是参数错误，比如是json格式错误
	zap.L().Error(err.Error())
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//"error": {
	//	"RegisterDTO.Name": "Name长度必须至少为3个字符",
	//	"RegisterDTO.RePassword": "RePassword必须等于Password"
	//}
	zap.L().Error("errors", zap.Any("errors", removeTopStruct(errs.Translate(trans))))
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(trans)),
	})
}

//func main() {
//	if err := InitTranslator("zh"); err != nil {
//		fmt.Println("zh")
//		return
//	}
//	router := gin.Default()
//	router.POST("/login", func(c *gin.Context) {
//		var loginDTO LoginDTO
//		if err := c.ShouldBind(&loginDTO); err != nil {
//			returnErrorMsg(c, err, trans)
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{
//			"msg": "login success",
//		})
//	})
//
//	router.POST("/register", func(c *gin.Context) {
//		var registerDTO RegisterDTO
//		if err := c.ShouldBind(&registerDTO); err != nil {
//			returnErrorMsg(c, err, trans)
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{
//			"msg": "register success",
//		})
//	})
//
//	router.Run(":8089")
//}
