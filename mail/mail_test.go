package mail

import "testing"

func TestSendMail(t *testing.T) {

	from := "xxxx@qq.com"
	to := "xxxxx@qq.com"
	subject := "Hello"
	body := "mail body <b>MMM<b>"

	m := Mail{
		From:    from,
		To:      []string{to},
		Subject: subject,
		Body:    body,
	}

	//m.SetMailType("text/plain") //加了这个html 就不会转码成html
	m.SendMail()
}
