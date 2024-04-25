package exEncrypt

import (
	"encoding/hex"
	"github.com/forgoer/openssl"
)

func DesEcbEncrypt(data, key string) (string, error) {
	encrypt, err := openssl.Des3ECBEncrypt([]byte(data), []byte(key), openssl.PKCS5_PADDING)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encrypt), nil
}

func DesEcbDecrypt(encryptData, key string) (string, error) {
	encrypt, err := hex.DecodeString(encryptData)
	if err != nil {
		return "", err
	}
	data, err1 := openssl.Des3ECBDecrypt(encrypt, []byte(key), openssl.PKCS5_PADDING)
	if err1 != nil {
		return "", err1
	}
	return string(data), nil
}
