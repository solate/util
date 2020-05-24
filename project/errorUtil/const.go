package errorUtil

// 错误码为6位数字，前三位为http类型，后三位为具体类型
var (

	// 400 请求参数类
	ErrParam    = New(400000, "request params error") // 请求参数异常
	ErrNotFound = New(404000, "not found")            // 未找到

	// 500 服务器类
	ErrSystem  = New(500000, "system error")       // 系统异常
	ErrUnknown = New(500001, "unknown")            // 未知错误
	ErrConfig  = New(500002, "parse config error") // 解析配置文件失败
)
