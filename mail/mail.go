package mail

import (
	"errors"
	"gopkg.in/gomail.v2"
)

// Conf ...
type (
	Conf struct {
		// 发件人密码，QQ邮箱这里配置授权码
		SPassword string
		// SMTP 服务器地址， QQ邮箱是smtp.qq.com
		SMTPAddr string
		// SMTP端口 QQ邮箱是25
		SMTPPort int
	}
	Mail struct {
		Conf    map[string]*Conf
		Message *gomail.Message
	}
)

func NewMail() *Mail {
	m := new(Mail)
	m.Conf = map[string]*Conf{
		"qq": {
			SMTPAddr:  `smtp.qq.com`,
			SPassword: `vljotearquwgbcia`,
			SMTPPort:  465,
		},
	}
	m.Message = gomail.NewMessage()
	return m
}

// SetConf 设置配置
func (t *Mail) SetConf(mailType string, conf *Conf) *Mail {
	if mailType == "" {
		return t
	}
	t.Conf[mailType] = conf
	return t
}

func (t *Mail) SetMessage(sender, title, body, attach string, recipientList []string) *Mail {
	t.Message.SetHeader("From", sender)
	t.Message.SetHeader("To", recipientList...)
	t.Message.SetHeader("Subject", title)
	t.Message.SetBody("text/html", body)
	if attach != "" {
		t.Message.Attach(attach)
	}
	return t
}

func (t *Mail) GetHeader(key string) string {
	return t.Message.GetHeader(key)[0]
}

func (t *Mail) Send(mailType string) error {
	if mailType == "" {
		mailType = "qq"
	}
	conf, ok := t.Conf[mailType]
	if !ok {
		return errors.New("conf is nil, mailType: " + mailType)
	}
	err := gomail.NewDialer(conf.SMTPAddr, conf.SMTPPort, t.GetHeader("From"), conf.SPassword).DialAndSend(t.Message)
	if err != nil {
		return err
	}
	return nil
}

func FormatBody(errPath, errLine, errMsg string) string {
	content := ""
	if errMsg == "" {
		return content
	}
	content += "FilePath: " + errPath + " Line: " + errLine + "<br />"
	content += "Error: " + errMsg + "<br />"
	return content
}
