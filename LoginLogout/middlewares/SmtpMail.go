package middlewares

import (
	"fmt"
	"jwt-authentication-golang/util"
	"log"
	"net/smtp"
	"strconv"
)

func SmtpMail(toEmail string) int {
	// Thông tin tài khoản email
	email := "kienvtinb@gmail.com"
	password := "Abc123456789@"

	// Thông tin người nhận
	to := []string{toEmail}

	// Cấu hình máy chủ SMTP
	smtpServer := "smtp.gmail.com"
	smtpPort := "25"

	random := util.Random()

	// Tạo thư gửi
	message := []byte("To: " + to[0] + "\r\n" +
		"Subject: Verification Codes\r\n" +
		"\r\n" + strconv.Itoa(random))

	// Xác thực tài khoản email
	auth := smtp.PlainAuth("", email, password, smtpServer)

	// Kết nối đến máy chủ SMTP
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, email, to, message)
	if err != nil {
		log.Println("Error sending email:", err)
		return -1
	}
	fmt.Println("Email sent!")

	return random
}
