package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// pprof 监听端口
var host string = "0.0.0.0:10000"

func init() {
	go func() {
		log.Println(http.ListenAndServe(host, nil))
	}()
}
