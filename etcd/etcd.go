package etcd

import (
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

// etcd 属性
type Etcd struct {
	Endpoints   []string         // etcd 集群列表
	DialTimeout int              // 连接超时时间，单位秒
	Username    string           // 用户名
	Password    string           // 密码
	Client      *clientv3.Client // 使用客户端

}

func NewEtcd(Endpoints []string, timeout int, username, password string) *Etcd {
	return &Etcd{Endpoints: Endpoints, DialTimeout: timeout, Username: username, Password: password}
}

// 初始化
func (s *Etcd) Init() (err error) {
	//初始化配置
	config := clientv3.Config{
		Endpoints: s.Endpoints, // 集群地址
		//TLS: conf,
		DialTimeout: time.Duration(s.DialTimeout) * time.Second, // 连接超时, 直接乘以秒，写时间容易搞错
		LogConfig: &zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.ErrorLevel),
			Development: true,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		},
		Username: s.Username,
		Password: s.Password,
	}
	// 建立连接
	s.Client, err = clientv3.New(config)
	return
}

// new etcd kv
func (s *Etcd) NewKv() clientv3.KV {
	return clientv3.NewKV(s.Client)
}

// new etcd lease
func (s *Etcd) NewLease() clientv3.Lease {
	return clientv3.NewLease(s.Client)
}

// new etcd watcher
func (s *Etcd) NewWatcher() clientv3.Watcher {
	return clientv3.NewWatcher(s.Client)
}
