package reply

import (
	"encoding/json"
)

type response struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 错误提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

const (
	ErrSuccessCode    = 200000
	ErrSuccessMessage = "success"
)

// json response
func Response(options ...func(*response)) (res []byte, err error) {
	// 设置参数
	response := &response{Code: ErrSuccessCode, Message: ErrSuccessMessage}
	for _, option := range options {
		option(response)
	}
	return json.Marshal(response)
}

// set err code and message,
func SetCodeAndMessage(code int, message string) func(*response) {
	return func(response *response) {
		response.Code = code
		response.Message = message
	}
}

// set data
func SetData(data interface{}) func(*response) {
	return func(response *response) {
		response.Data = data
	}
}

// set err code
func SetCode(code int) func(*response) {
	return func(response *response) {
		response.Code = code
	}
}

// set err message
func SetMessage(message string) func(*response) {
	return func(response *response) {
		response.Message = message
	}
}
