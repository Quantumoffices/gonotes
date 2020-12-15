package main

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

//163邮箱-服务器地址：
//POP3服务器: pop.163.com
//SMTP服务器: smtp.163.com
//IMAP服务器: imap.163.com
func main() {
	e := email.NewEmail()
	//e.From = "pingyeaa@163.com"
	e.From = "平也 <mail1217105106@163.com>"
	e.To = []string{"1217105106@qq.com"}
	e.Subject = "发现惊天大秘密！"
	e.Text = []byte("平也好帅好有智慧哦~")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "mail1217105106@163.com", "xzy1883182019", "smtp.163.com"))
	if err != nil {
		panic(err)
	}
}
