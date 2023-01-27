package common

type CustomError struct {
	ErrorCode Code   `json:"code"`
	ErrorMsg  string `json:"msg"`
}

func (c *CustomError) Error() string {
	return c.ErrorMsg
}

func NewCustomError(code Code) CustomError {
	return CustomError{
		ErrorCode: code,
		ErrorMsg:  ToMsg(code),
	}
}
