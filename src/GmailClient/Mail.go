package GmailClient

import (
	"net/smtp"
	"bytes"
	"encoding/base64"
	"strings"
	"os"
)

type Mail struct {
	To string
	Body string
	Subject string
}

type MailClientConfig struct {
	SmtpHost string
	SmtpServer string
	Username string
	Password string
}

func (config MailClientConfig) getFrom() string {
	return config.Username + "@gmail.com"
}

func (config MailClientConfig) getAuth() smtp.Auth {
	 auth := smtp.PlainAuth("", config.Username, config.Password, config.SmtpHost)
	 return auth
}

func newGmailConfig() MailClientConfig {
	config := MailClientConfig{}
	config.SmtpHost = os.Getenv("GMAIL_SMTP_HOST")
	config.Username = os.Getenv("GMAIL_USERNAME")
	config.Password = os.Getenv("GMAIL_PASSWORD")
	config.SmtpServer = os.Getenv("GMAIL_SMTP_SERVER")
	return config
}

func send(m Mail, config MailClientConfig) error {
		auth := config.getAuth()
		body := encodeSubject(m.Subject) + m.Body

		err := smtp.SendMail(config.SmtpServer, auth, config.getFrom(), []string{m.To}, []byte(body))

		if err != nil {
			return err
		}
    return nil
}

// サブジェクトを MIME エンコードする
func encodeSubject(subject string) string {
    // UTF8 文字列を指定文字数で分割する
    b := bytes.NewBuffer([]byte(""))
    strs := []string{}
    length := 13
    for k, c := range strings.Split(subject, "") {
        b.WriteString(c)
        if k%length == length-1 {
            strs = append(strs, b.String())
            b.Reset()
        }
    }
    if b.Len() > 0 {
        strs = append(strs, b.String())
    }
    // MIME エンコードする
    b2 := bytes.NewBuffer([]byte(""))
    b2.WriteString("Subject:")
    for _, line := range strs {
        b2.WriteString(" =?utf-8?B?")
        b2.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
        b2.WriteString("?=\r\n")
    }
    return b2.String()
}