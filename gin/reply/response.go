package reply

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 服务回复信息
type Response struct {
	ErrCode int         `json:"err_code"` // 错误码
	ErrMsg  string      `json:"err_msg"`  // 错误提示信息
	Data    interface{} `json:"data"`     // 返回数据
}

const (
	ErrSuccessCode = 200000 //6位数错误码
	ErrSuccessMsg  = "success"
)

// 返回成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{ErrSuccessCode, ErrSuccessMsg, data})
}

// 返回失败, 但是参数可以用方法设置
//options: httpCode, errCode int, errMsg string, data
func Error(c *gin.Context, httpCode int, options ...func(*Response)) {

	// 设置参数
	res := &Response{}
	for _, option := range options {
		option(res)
	}
	c.JSON(httpCode, res)
}

// 设置错误码, 一般有code 就会有msg
func SetErrCodeAndMsg(code int, msg string) func(*Response) {
	return func(response *Response) {
		response.ErrCode = code
		response.ErrMsg = msg
	}
}

// 设置数据
func SetData(data interface{}) func(*Response) {
	return func(response *Response) {
		response.Data = data
	}
}

// 设置错误码
func SetErrCode(code int) func(*Response) {
	return func(response *Response) {
		response.ErrCode = code
	}
}

// 设置错误码
func SetErrMsg(msg string) func(*Response) {
	return func(response *Response) {
		response.ErrMsg = msg
	}
}
