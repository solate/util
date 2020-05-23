# util
go util  封装一些帮助包, 方便引入使用

## color 

改变打印字体颜色

## cron

cron 表达式字符串处理使用封装， 这个源码需要修改，后续自己写一个

## etcd 

etcd v3 客户端封装，项目直接使用单例


## gin

### reply 

统一返回值封装

返回json串, 包含错误码，错误信息，数据三部分(err_code, err_msg, data)

## go

go 原生封装

### copy

两个简单的struct直接进行拷贝

## gofile

对file 操作封装，少些代码

## gostring

对string 的一些封装

## gotemplate

对模板的封装

## network

网络相关的封装，现在只有获取本机IP


## orm 

### gorm 

对gorm mysql 使用封装，返回带连接池的db, 不用每次copy同样的东西


## log

### zap

对高性能zap库的封装


## mail

使用gomail 发送邮件包