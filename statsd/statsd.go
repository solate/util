package statsd

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

const (
	// UDP packet limit, see
	// https://en.wikipedia.org/wiki/User_Datagram_Protocol#Packet_structure
	UDP_MAX_PACKET_SIZE int = 64 * 1024
)

type Statsd struct {
	protocol string   //协议
	address  string   //地址
	rate     float64  //频率数
	Conn     net.Conn //连接
}

//新建实例
func NewStatsd(protocol, address string, rate float64) *Statsd {
	return &Statsd{protocol: protocol, address: address, rate: rate}
}

//初始化
func (s *Statsd) Init() (err error) {
	connection, err := net.Dial(s.protocol, s.address)
	if err != nil {
		return
	}

	s.Conn = connection
	return
}

//写入telegraf-statsd
func (s *Statsd) StatsdPush(measurement string, tagSet map[string]string, fieldValue string, suffix string) {
	w := bufio.NewWriter(s.Conn)
	statsDStr := s.statsdLine(measurement, tagSet, fieldValue, suffix)
	w.Write([]byte(statsDStr))
	w.Flush()
}

func (s *Statsd) statsdLine(measurement string, tagSet map[string]string, fieldValue string, suffix string) string {
	var builder strings.Builder
	builder.WriteString(measurement)
	builder.WriteString(":")
	builder.WriteString(fieldValue)
	builder.WriteString(suffix)

	//拼接tag
	if len(tagSet) > 0 {
		builder.WriteString("|#")
		joint := ""
		for tagK, tagV := range tagSet {
			if joint != "" {
				builder.WriteString(joint)
			}
			builder.WriteString(tagK)
			builder.WriteString(":")
			builder.WriteString(tagV)
			joint = ","
		}
	}

	//拼接rate
	if s.rate > 0 && s.rate < 1 {
		builder.WriteString("|@")
		builder.WriteString(strconv.FormatFloat(s.rate, 'E', -1, 32))
	}

	builder.WriteString("\n") //每一行打一个回车，telegraf 是使用回车进行读取的
	return builder.String()
}
