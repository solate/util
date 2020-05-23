package orm

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	maxIdleConns    = 10                //连接池中最大空闲连接数
	maxOpenConns    = 20                //打开的最大连接数
	connMaxLifetime = 300 * time.Second //连接的最大空闲时间(可选)
	dialect         = "mysql"
)

type Orm struct {
	Host     string // 数据库地址, localhost
	Port     string // 数据库端口号 3306
	Database string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
	*Options        // 可选参数
	*gorm.DB        // DB
}

// options
type Options struct {
	MaxIdleConns    int           // 连接池中最大空闲连接数
	MaxOpenConns    int           // 打开的最大连接数
	ConnMaxLifetime time.Duration // 连接的最大空闲时间, 单位秒
	Debug           bool          // 连接是否为debug, debug 会开启sql 打印
}

func NewOrm(host, prot, database, username, password string) (o *Orm) {
	return &Orm{Host: host, Port: prot, Database: database, Username: username, Password: password}
}

func (s *Orm) SetOption(maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration, debug bool) *Orm {
	s.Options = &Options{
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
		ConnMaxLifetime: connMaxLifetime,
		Debug:           debug,
	}
	return s
}

// init mysql client
func (s *Orm) Init() (db *gorm.DB, err error) {

	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
	)

	db, err = gorm.Open(dialect, dsn)
	if err != nil {
		return
	}

	db.LogMode(s.Options.Debug) //是否开启sql打印

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
