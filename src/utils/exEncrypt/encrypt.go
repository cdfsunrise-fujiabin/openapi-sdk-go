package exEncrypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type CdfEncrypt interface {
	Encrypt(sign Sign) string
}

type Md5Encrypt struct {
}

func NewMd5Encrypt() *Md5Encrypt {
	return &Md5Encrypt{}
}

func (e *Md5Encrypt) Encrypt(sign Sign) (string, error) {
	sign.collect()
	plainTxt, err := sign.GenSign()
	if err != nil {
		return "", err
	}
	return e.Md5(plainTxt), nil
}

func (e *Md5Encrypt) Md5(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func (e *Md5Encrypt) UpperEncrypt(sign Sign) (string, error) {
	sign.collect()
	plainTxt, err := sign.GenSign()
	if err != nil {
		return "", err
	}
	return e.UpperMd5(plainTxt), nil
}

func (e *Md5Encrypt) UpperMd5(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return strings.ToUpper(hex.EncodeToString(md5Hash[:])) // by referring to it as a string
}
