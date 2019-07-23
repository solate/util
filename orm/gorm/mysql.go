package gorm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Mysql struct {
	Host     string // 数据库地址, localhost
	Port     string // 数据库端口号 3306
	Database string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
	*Options        // 可选参数
}

// 生成dsn 字符串
func (s *Mysql) args() (dialect, dsn string) {
	dialect = "mysql"

	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
	)
	return
}

// 初始化一个mysql
func InitMysql(mysql *Mysql) (db *gorm.DB, err error) {

	if mysql.Options == nil {
		mysql.Options = &Options{}
	}

	client := &params{
		args:    mysql,
		Options: mysql.Options,
	}
	return client.new()
}
