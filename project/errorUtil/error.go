package errorUtil

import (
	"fmt"
	"runtime"
	"strconv"
)

type Error struct {
	Code    int    // 错误码
	Message string // 错误信息
	Trace   string // 错误调用队列
}

func New(code int, message string) *Error {
	_, file, line, _ := runtime.Caller(1)
	trace := file + ":" + strconv.Itoa(line)
	return &Error{Code: code, Message: message, Trace: trace}
}

func (s Error) Error() string {
	return fmt.Sprintf("code = [%v]; message = [%s]; trace = [%s]", s.Code, s.Message, s.Trace)
}
