package common

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
)

func MailSendCode(mail, code string) error {
	config := viper.New()
	//在项目中查找配置文件的路径，可以使用相对路径，也可以使用绝对路径
	config.AddConfigPath("./config")
	//配置文件名（不带扩展名）
	config.SetConfigName("application")
	//设置文件类型，这里是yaml文件
	config.SetConfigType("yml")
	//查找并读取配置文件
	err := config.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	username := config.GetString("mail.username")
	password := config.GetString("mail.password")

	e := email.NewEmail()
	e.From = "Get <" + username + ">"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为:<h1>" + code + "</h1>")
	err = e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", username, password, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}
