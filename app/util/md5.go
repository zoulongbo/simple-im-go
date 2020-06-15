package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//返回小写
func Md5Encode(data string) string{
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)

	return  hex.EncodeToString(cipherStr)

}

//大写
func MD5Encode(data string) string{
	return strings.ToUpper(Md5Encode(data))
}
//密码校验
func ValidatePassword(plainPwd,salt,password string) bool{
	return Md5Encode(plainPwd+salt)==password
}

//生成密码
func MakePassword(plainPwd,salt string) string{
	return Md5Encode(plainPwd+salt)
}