package common

type Code int

const (
	CodeSuccess = 1000 + iota
	CodeServerError
	CodeInvalidPassword
	CodeUserExist
	CodeVoteTimeExpired
)

var MsgMap = map[Code]string{
	CodeSuccess:         "success",
	CodeServerError:     "服务器异常",
	CodeInvalidPassword: "用户名或密码错误",
	CodeUserExist:       "用户已存在",
	CodeVoteTimeExpired: "投票时间已过期",
}

func ToMsg(code Code) string {
	msg, ok := MsgMap[code]
	if !ok {
		return MsgMap[CodeServerError]
	}
	return msg
}
