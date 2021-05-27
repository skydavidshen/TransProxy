package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	rand2 "math/rand"
)

func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand2.Intn(len(bytes))])
	}
	return string(result)
}

func GenUUID() string {
	id, _ := gonanoid.New(32)
	return id
}