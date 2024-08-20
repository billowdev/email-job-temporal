package helpers

import (
	"bytes"
	"crypto/tls"
	"errors"
	"strings"

	"fmt"
	"net"
	"net/smtp"
	"text/template"
	"time"

	"github.com/billowdev/email-job-temporal/pkg/configs"
)

func sendEmailHelper(conn *smtp.Client, from, to, msg string) error {
	if err := conn.Mail(from); err != nil {
		return err
	}
	if err := conn.Rcpt(to); err != nil {
		return err
	}

	wc, err := conn.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = wc.Write([]byte(msg))
	if err != nil {
		return err
	}

	err = wc.Close()
	if err != nil {
		return err
	}

	return nil
}

// https://gist.github.com/chrisgillis/10888032?permalink_comment_id=2553469
func ConnectSecureSMTP(host, port, username, password string) (*smtp.Client, error) {
	servername := fmt.Sprintf("%v:%v", host, port)
	hostPort, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", username, password, hostPort)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: configs.SMTP_INSECURE_SKIP_VERIFY,
		ServerName:         host,
	}

	// TODO: set timeout
	dialer := &net.Dialer{
		Timeout: 15 * time.Second, // 30-second timeout
	}

	// Dial and establish TLS connection
	conn, err := tls.DialWithDialer(dialer, "tcp", servername, tlsconfig)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println(err.Error())
			fmt.Println("Connection timed out")
			errInfo := errors.New("connection timed out")
			// return nil, errInfo
			_ = errInfo
			return nil, err
		} else {
			fmt.Println("Error connecting to SMTP server:", err)
			fmt.Println(err.Error())
			errInfo := errors.New("error connecting to SMTP server")
			// return nil, errInfo
			_ = errInfo
			return nil, err
		}
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}

	if configs.SMTP_START_TLS {
		if err := c.StartTLS(tlsconfig); err != nil {
			return nil, err
		}
	}
	if configs.SMTP_IS_AUTH_REQUIRED {
		// Auth
		if err = c.Auth(auth); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func ConnectSimpleSMTP(host, port, username, password string) (*smtp.Client, error) {
	servername := fmt.Sprintf("%v:%v", host, port)
	// Dial the SMTP server with the timeout
	c, err := smtp.Dial(servername)
	if err != nil {
		fmt.Println("Failed to connect to SMTP server:", err)
		return nil, err
	}

	return c, nil
}

func SendEmail(from, to, subject, msg string, ccData []string) error {
	host := configs.SMTP_HOST
	port := configs.SMTP_PORT
	username := configs.SMTP_USERNAME
	password := configs.SMTP_PASSWORD

	// c, err := ConnectSMTP(host, port, username, password)
	var err error
	var c *smtp.Client

	// c, err = ConnectSimpleSMTP(host, port, username, password)
	// if err != nil {
	// 	return err
	// }
	// defer c.Close()

	c, err = ConnectSecureSMTP(host, port, username, password)
	if err != nil {
		return err
	}
	defer c.Close()

	// Setup headers
	cc := strings.Join(ccData, ",")

	headers := map[string]string{
		"From":         from,
		"To":           to,
		"CC":           cc,
		"Subject":      subject,
		"MIME-version": "1.0",
		"Content-Type": "text/html; charset=\"UTF-8\"",
	}
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + msg

	err = sendEmailHelper(c, from, to, string(message))
	if err != nil {
		return err
	}

	return err
}

func ParseHTMLTemplateHelper[T interface{}](htmlTemplate string, htmlArgs T) (string, error) {
	t, err := template.New("email_template").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return "", err
	}

	var tplBuffer bytes.Buffer
	err = t.Execute(&tplBuffer, htmlArgs)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return "", err
	}

	newHTMLData := tplBuffer.String()

	return newHTMLData, nil
}
