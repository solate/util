package mail

import "gopkg.in/gomail.v2"

const (
	ServerHost = "smtp.qq.com" // 邮件服务器
	Port       = 465           // SMTP 协议默认使用25, qq使用465
	User       = "xxxx@qq.com" // 发送者
	Password   = ""            // 密码

)

type Mail struct {
	From     string   // 发送人
	To       []string // 发送给谁
	Subject  string   // 标题
	Body     string   // 主体内容
	Attach   string   // 附件文件路径
	MailType string   //文件类型
}

// 设置发送邮件的邮件类型
// text/html & text/plain
func (s *Mail) SetMailType(t string) {
	s.MailType = t
}

// 发送邮件
// form 发送人
// to 发送给谁
// subject 标题
// body 主体内容
// attach 附件文件地址
func (s *Mail) SendMail() (err error) {

	var mailType string
	if s.MailType == "" {
		mailType = "text/html"
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.From)
	m.SetHeader("To", s.To...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", s.Subject)
	m.SetBody(mailType, s.Body)
	if s.Attach != "" { //附件
		m.Attach(s.Attach)
	}

	d := gomail.NewDialer(ServerHost, Port, User, Password)
	return d.DialAndSend(m)
}
