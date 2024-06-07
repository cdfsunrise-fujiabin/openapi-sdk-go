package exEncrypt

import (
	"fmt"
	"sort"
	"strings"
)

type Sign interface {
	// Collect 收集阶段
	collect()

	// prepare 准备阶段
	prepare() bool

	// GenSign 签名阶段
	GenSign() (string, error)
}

type BasicSign struct {
}

func NewBasicSign() *BasicSign {
	return &BasicSign{}
}

func (s *BasicSign) collect() {

}

func (s *BasicSign) prepare() bool {

	return true
}

func (s *BasicSign) GenSign() (string, error) {
	s.prepare()
	return "result", nil
}

func sortStrings(strs []string) []string {
	sort.Strings(strs)
	return strs
}

type CdfSign struct {
	key       string
	paramKeys []string
	sortKeys  []string
	params    map[string]string
}

func NewCdfSign(key string, params map[string]string) *CdfSign {
	return &CdfSign{
		key:    key,
		params: params,
	}
}

func (s *CdfSign) Collect() {
	s.params["key"] = s.key
	var paramKeys []string
	for k, _ := range s.params {
		paramKeys = append(paramKeys, k)
	}
	s.paramKeys = paramKeys
}

func (s *CdfSign) prepare() bool {
	s.sortKeys = sortStrings(s.paramKeys)
	return true
}

func (s *CdfSign) GenSign() (string, error) {
	s.Collect()
	s.prepare()
	var plaintexts []string
	for _, key := range s.sortKeys {
		if s.params[key] != "" {
			plaintexts = append(plaintexts, fmt.Sprintf("%s=%s", key, s.params[key]))
		}
	}
	joined := strings.Join(plaintexts, "&")
	encrypt := NewMd5Encrypt()
	sign := encrypt.Md5(joined)
	return sign, nil
}
