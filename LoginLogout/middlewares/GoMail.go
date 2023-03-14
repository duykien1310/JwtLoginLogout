package middlewares

import (
	"fmt"
	"jwt-authentication-golang/util"
	"strconv"

	"gopkg.in/gomail.v2"
)

func GoMail(Tomail string) int {
	// Tạo một đối tượng Message mới
	m := gomail.NewMessage()

	// Đặt giá trị từ, đến, chủ đề, nội dung của email
	m.SetHeader("From", "kienvtinb@gmail.com")
	m.SetHeader("To", Tomail)
	m.SetHeader("Subject", "Veritification Codes")
	random := util.Random()
	m.SetBody("text/plain", "Code:\n"+strconv.Itoa(random))

	// Tạo một đối tượng Dialer mới để kết nối đến máy chủ SMTP
	d := gomail.NewDialer("smtp.gmail.com", 587, "kienvtinb@gmail.com", "Abc@123456789")

	// Thực hiện xác thực bằng tên người dùng và mật khẩu
	// d.TLSConfig.InsecureSkipVerify = true

	// Gửi email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent!")
	}
	return random
}
