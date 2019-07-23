package gorm

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	maxIdleConns    = 10                //连接池中最大空闲连接数
	maxOpenConns    = 20                //打开的最大连接数
	connMaxLifetime = 300 * time.Second //连接的最大空闲时间(可选)
)

type args interface {
	args() (dialect string, args string)
}

type params struct {
	args
	*Options
}

// 可选参数
type Options struct {
	MaxIdleConns    int           // 连接池中最大空闲连接数
	MaxOpenConns    int           // 打开的最大连接数
	ConnMaxLifetime time.Duration // 连接的最大空闲时间, 单位秒
	IsDebug         bool          // 连接是否为debug, debug 会开启sql 打印
}

//初始化数据库客户端
func (s *params) new() (db *gorm.DB, err error) {

	db, err = gorm.Open(s.args.args())
	if err != nil {
		return
	}

	db.LogMode(s.IsDebug) //是否开启sql打印

	// pool  连接池
	//默认连接池连接数
	if s.MaxIdleConns == 0 {
		s.MaxIdleConns = maxIdleConns
	}
	if s.MaxOpenConns == 0 {
		s.MaxOpenConns = maxOpenConns
	}
	if s.ConnMaxLifetime == 0 {
		s.ConnMaxLifetime = connMaxLifetime
	}
	db.DB().SetMaxIdleConns(s.MaxIdleConns)       //连接池中最大空闲连接数
	db.DB().SetMaxOpenConns(s.MaxOpenConns)       //打开的最大连接数
	db.DB().SetConnMaxLifetime(s.ConnMaxLifetime) //连接的最大空闲时间(可选)

	return
}
