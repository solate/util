package reply

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 服务回复结构
type response struct {
	HttpCode int `json:"http_code"` // 响应服务http状态码
	body
}

// 回复body结构
type body struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 错误提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

const (
	ErrSuccessCode = 200000 //6位数错误码
	ErrSuccessMsg  = "success"
)

func Response(c *gin.Context, options ...func(*response)) {
	// 设置参数
	res := &response{http.StatusOK, body{Code: ErrSuccessCode, Message: ErrSuccessMsg}}
	for _, option := range options {
		option(res)
	}
	c.JSON(res.HttpCode, res.body)
}

// 设置数据
func SetHttpCode(code int) func(*response) {
	return func(response *response) {
		response.HttpCode = code
	}
}

// 设置错误码, 一般有code 就会有message
func SetCodeMessage(code int, message string) func(*response) {
	return func(response *response) {
		response.Code = code
		response.Message = message
	}
}

// 设置数据
func SetData(data interface{}) func(*response) {
	return func(response *response) {
		response.Data = data
	}
}

// 设置错误码
func SetCode(code int) func(*response) {
	return func(response *response) {
		response.Code = code
	}
}

// 设置错误码
func SetMessage(message string) func(*response) {
	return func(response *response) {
		response.Message = message
	}
}
