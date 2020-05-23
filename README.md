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

### gofile

对file 操作封装，少些代码

### gostring

对string 的一些封装

### gotemplate

对模板的封装

### network

网络相关的封装，现在只有获取本机IP

## log

封装了zap logger, 单例使用

## mail

使用gomail 发送邮件包

## orm 

对gorm mysql 使用封装，返回带连接池的db, 不用每次copy同样的东西

## project

生成项目，需要用到的工具包

### configuration

读取toml格式的配置文件，并解析成结构体

### pprof

默认使用10000端口开启pprof 监听，只要在import中引用就好

### reply 

普通返回json格式内容

## retry

重试方法

## statsd 

statsd udp 发送消息到influxdb中

## uuid

google uuid 生成

