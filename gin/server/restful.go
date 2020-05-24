package server

import "github.com/gin-gonic/gin"

type Server struct {
	port string // http 接口
}

func New(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start(f func(r *gin.Engine)) (err error) {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	f(r) //注册路由

	//启动
	r.Run(":" + s.port)
	return
}
