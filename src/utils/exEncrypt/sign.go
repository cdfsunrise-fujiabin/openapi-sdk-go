package exEncrypt

import (
	"fmt"
	"github.com/pkg/errors"
	"sort"
	"strings"
)

type Sign interface {
	// Collect 收集阶段
	Collect()

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

func (s *BasicSign) Collect() {

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
	s.prepare()
	var plaintexts []string
	for _, key := range s.sortKeys {
		if s.params[key] != "" {
			plaintexts = append(plaintexts, fmt.Sprintf("%s=%s", key, s.params[key]))
		}
	}
	return strings.Join(plaintexts, "&"), nil
}

// client_secret+timestamp+client_id+username+password+grant_type+scope+client_secret
type ZltAuthParam struct {
	ClientSecret string
	TimeStamp    string
	ClientId     string
	Username     string
	Password     string
	GrantType    string
	Scope        string
}

func NewZltAuthSign(data ZltAuthParam) *ZltAuthSign {
	return &ZltAuthSign{
		Data: data,
	}
}

type ZltAuthSign struct {
	Data ZltAuthParam
}

func (s *ZltAuthSign) Collect() {
}

func (s *ZltAuthSign) prepare() bool {
	if s.Data.ClientId == "" || s.Data.ClientSecret == "" || s.Data.TimeStamp == "" || s.Data.Username == "" || s.Data.Password == "" || s.Data.GrantType == "" || s.Data.Scope == "" {
		return false
	}
	return true
}

// client_secret+timestamp+client_id+username+password+grant_type+scope+client_secret 加密顺序
func (s *ZltAuthSign) GenSign() (string, error) {
	if !s.prepare() {
		return "", errors.New("缺少数据生成失败")
	}
	plaintexts := s.Data.ClientSecret + s.Data.TimeStamp + s.Data.ClientId + s.Data.Username + s.Data.Password + s.Data.GrantType + s.Data.Scope + s.Data.ClientSecret
	return plaintexts, nil
}

type ZltDataSign struct {
	ClientSecret string
	paramKeys    []string
	sortKeys     []string
	params       map[string]string
}

func NewZltDataSign(ClientSecret string, params map[string]string) *ZltDataSign {
	return &ZltDataSign{
		ClientSecret: ClientSecret,
		params:       params,
	}
}
func (s *ZltDataSign) Collect() {
	var paramKeys []string
	for k, _ := range s.params {
		paramKeys = append(paramKeys, k)
	}
	s.paramKeys = paramKeys
}

func (s *ZltDataSign) prepare() bool {
	s.sortKeys = sortStrings(s.paramKeys)
	return true
}

func (s *ZltDataSign) GenSign() (string, error) {
	s.prepare()
	var plaintexts []string
	for _, key := range s.sortKeys {
		if s.params[key] != "" {
			plaintexts = append(plaintexts, fmt.Sprintf("%s=%s", key, s.params[key]))
		}
	}
	text := s.ClientSecret + strings.Join(plaintexts, "&") + s.ClientSecret
	return text, nil
}
