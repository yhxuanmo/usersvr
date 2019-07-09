package utils

import (
	"math/rand"
	"time"
)

func GenVerifyCode() string {
	str := "0123456789"
	var res []byte
	byteStr := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		res = append(res, byteStr[r.Intn(len(byteStr))])
	}
	return string(res)
}
