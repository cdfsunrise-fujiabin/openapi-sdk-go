package httpClient

import (
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exEncrypt"
)

type CdfSunriseRequestClient struct {
	host string
	//appId     string
	//appKey    string
	rsaPubKey     string
	rsaPrivateKey string
	xrsa          *exEncrypt.XRsa
}

func NewCdfSunriseRequestClient(host, rsaPubKey, rsaPrivateKey string) *CdfSunriseRequestClient {
	xrsa, _ := exEncrypt.NewXRsa([]byte(rsaPubKey), []byte(rsaPrivateKey))
	return &CdfSunriseRequestClient{
		host: host,
		//appId:     appId,
		//appKey:    appKey,
		rsaPubKey:     rsaPubKey,
		rsaPrivateKey: rsaPrivateKey,
		xrsa:          xrsa,
	}
}

var _headers = map[string]string{
	"content-type": "application/json",
}

func GenHeaders(customHeaders map[string]string) map[string]string {
	h := make(map[string]string)
	for k, v := range _headers {
		h[k] = v
	}
	for k, v := range customHeaders {
		h[k] = v
	}
	return h
}

func (t *CdfSunriseRequestClient) EncryptByRsa(raw string) string {
	if raw == "" || t.xrsa == nil {
		return ""
	}
	encrypt, err := t.xrsa.PublicEncrypt(raw)
	if err != nil {
		return ""
	}
	return encrypt
}

func (t *CdfSunriseRequestClient) DecryptByRsa(encryptContent string) string {
	if encryptContent == "" || t.xrsa == nil {
		return ""
	}
	rawData, err := t.xrsa.PrivateDecrypt(encryptContent)
	if err != nil {
		return ""
	}
	return rawData
}
