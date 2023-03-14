package util

import (
	"math/rand"
	"time"
)

func Random() int {
	rand.Seed(time.Now().UnixNano()) // Đặt giá trị seed ngẫu nhiên dựa trên thời gian hiện tại

    // Tạo số ngẫu nhiên gồm 6 chữ số
    min := 100000
    max := 999999
    randomNum := rand.Intn(max-min+1) + min
    
	return randomNum
}