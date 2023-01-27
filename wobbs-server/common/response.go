package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func response(ctx *gin.Context, httpStatus int, resp Response) {
	ctx.JSON(httpStatus, resp)
}

func FailByMsg(ctx *gin.Context, msg string) {
	response(ctx, http.StatusOK, Response{
		Code: CodeServerError,
		Msg:  msg,
		Data: nil,
	})
}

func FailByCode(ctx *gin.Context, code Code) {
	response(ctx, http.StatusOK, Response{
		Code: int(code),
		Msg:  ToMsg(code),
		Data: nil,
	})
}

func SuccessByMsg(ctx *gin.Context, msg string) {
	response(ctx, http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  msg,
		Data: nil,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	response(ctx, http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  ToMsg(CodeSuccess),
		Data: data,
	})
}
