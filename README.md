# util
go util  封装一些帮助包, 方便引入使用

## color 

改变打印字体颜色



## gin

### reply 

统一返回值封装

返回json串, 包含错误码，错误信息，数据三部分(err_code, err_msg, data)



## orm 

### gorm 

对gorm mysql 使用封装，返回带连接池的db, 不用每次copy同样的东西


## log

### zap

对高性能zap库的封装


## mail

使用gomail 发送邮件包