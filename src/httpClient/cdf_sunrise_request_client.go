package httpClient

type CdfSunriseRequestClient struct {
	host      string
	appId     string
	appKey    string
	rsaPubKey string
}

func NewCdfSunriseRequestClient(host, appId, appKey, rsaPubKey string) *CdfSunriseRequestClient {
	return &CdfSunriseRequestClient{
		host:      host,
		appId:     appId,
		appKey:    appKey,
		rsaPubKey: rsaPubKey,
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
