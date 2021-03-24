package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptData struct {
	Password string `json:"password"`
}

// 生成加密
func BcryptGenerate(password string) (string, error) {
	// 数据实例
	bcryptData := BcryptData{}
	bcryptData.Password = password
	// 加密操作
	hash, err := bcrypt.GenerateFromPassword([]byte(bcryptData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 验证加密
func BcryptCompare(hashedPwd string, inputPwd string) bool {
	// 数据实例
	bcryptData := BcryptData{}
	bcryptData.Password = hashedPwd
	// 加密操作
	err := bcrypt.CompareHashAndPassword([]byte(bcryptData.Password), []byte(inputPwd))
	if err != nil {
		return false
	}
	return true
}
