package email

import (
	"fmt"
	"log"
	"regexp"

	"github.com/go-gomail/gomail"
)

var defaultEmailerManager *Manager

type Manager struct {
	handler *gomail.Dialer
	ch      chan Information
	reg     *regexp.Regexp
}

func New() *Manager {
	m := &Manager{
		handler: gomail.NewDialer(
			C.Service,
			C.Port,
			C.Account,
			C.Password,
		),
		ch:  make(chan Information, 1024),
		reg: regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`),
	}
	return m
}

func (m *Manager) SendEmail(info Information) {
	m.ch <- info
}

func (m *Manager) Run() {
	for {
		info := <-m.ch
		m.sendEmail(info)
	}
}

func (m *Manager) sendEmail(info Information) {
	content := gomail.NewMessage()
	content.SetAddressHeader("From", C.Account, C.Name)
	content.SetHeader("To", info.To)
	content.SetHeader("Subject", "[CHAT]身份认证邮件")
	content.SetBody("text/html", fmt.Sprintf("您的验证码是：%s", info.Code))
	err := m.handler.DialAndSend(content)
	if err != nil {
		log.Printf("[EMAIL ERROR] %+v\n", err)
	}
}

func (m *Manager) IsEmail(account string) bool {
	return m.reg.MatchString(account)
}

func GetDefaultEmailerManager() *Manager {
	return defaultEmailerManager
}
