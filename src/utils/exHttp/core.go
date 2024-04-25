package exHttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	myLog "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type (
	HttpRequest struct {
		verbose     bool
		ctx         context.Context
		url         string
		requestBody string
		proxy       string
		headers     map[string]string
		timeout     time.Duration
	}

	HttpRequestOption func(r *HttpRequest)
)

func NewHttpRequest(ctx context.Context, host, path string, opts ...HttpRequestOption) *HttpRequest {
	r := &HttpRequest{
		ctx: ctx,
		url: fixHost(host) + fixRelativeUrl(path),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (h *HttpRequest) SwitchVerbose(verbose bool) {
	h.verbose = verbose
}

func (h *HttpRequest) Get() ([]byte, error) {
	return h.sendHttpRequest(http.MethodGet)
}

func (h *HttpRequest) GetUnmarshal() (any, error) {
	return h.sendHttpRequestUnmarshall(http.MethodGet)
}

func (h *HttpRequest) Post() ([]byte, error) {
	return h.sendHttpRequest(http.MethodPost)
}

func (h *HttpRequest) PostUnmarshall() (any, error) {
	return h.sendHttpRequestUnmarshall(http.MethodPost)
}

func (h *HttpRequest) Put() ([]byte, error) {
	return h.sendHttpRequest(http.MethodPut)
}

func (h *HttpRequest) PutUnmarshall() (any, error) {
	return h.sendHttpRequestUnmarshall(http.MethodPut)
}

func (h *HttpRequest) Delete() ([]byte, error) {
	return h.sendHttpRequest(http.MethodDelete)
}

func (h *HttpRequest) DeleteUnmarshall() (any, error) {
	return h.sendHttpRequestUnmarshall(http.MethodDelete)
}

func (h *HttpRequest) GetProxy() string {
	return h.proxy
}

func (h *HttpRequest) sendHttpRequest(method string) ([]byte, error) {
	t := time.Now()
	req, err := h.genHttpRequest(method, h.url, h.requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "genHttpRequest failed")
	}

	if h.verbose {
		dumpReq, er := httputil.DumpRequest(req, true)
		if er == nil {
			fmt.Println(string(dumpReq))
		}
	}

	respModel, err := h.doHttpRequest(req)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("doHttpRequest failed(%v)", h.proxy))
	}

	if time.Since(t).Milliseconds() > 3000 {
		myLog.WithContext(h.ctx).WithFields(myLog.Fields{"url": h.url, "proxy": h.proxy, "elipse": time.Since(t).Milliseconds()}).Info("http request slow")
	}

	return respModel, nil
}

func (h *HttpRequest) sendHttpRequestUnmarshall(method string) (any, error) {
	respModel, err := h.sendHttpRequest(method)
	if err != nil {
		return nil, err
	}

	var r any
	err = json.Unmarshal(respModel, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (h *HttpRequest) PutForm(params url.Values) ([]byte, error) {
	return h.requestByForm("PUT", params)
}

func (h *HttpRequest) PostForm(params url.Values) ([]byte, error) {
	return h.requestByForm("POST", params)
}

func (h *HttpRequest) doHttpRequest(req *http.Request) ([]byte, error) {
	var httpClient *http.Client
	if h.proxy != "" {
		urlI := url.URL{}
		urlProxy, err := urlI.Parse(h.proxy)

		if err == nil {
			httpClient = &http.Client{
				Transport: &http.Transport{
					Proxy:               http.ProxyURL(urlProxy),
					TLSHandshakeTimeout: 20 * time.Second,
					MaxIdleConnsPerHost: 10,
					MaxIdleConns:        200,
					IdleConnTimeout:     120 * time.Second,
				},
			}
		}
	} else {
		httpClient = http.DefaultClient
	}

	if h.timeout > 0 {
		httpClient.Timeout = h.timeout
	} else {
		httpClient.Timeout = time.Second * 15
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send http request failed")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "read data from http response failed")
	}

	return respBody, nil
}

func (h *HttpRequest) genHttpHeaders(req *http.Request) {
	req.Header.Add("Accept-Encoding", "identity")
	if h.headers != nil && len(h.headers) > 0 {
		for k, v := range h.headers {
			req.Header[k] = []string{v}
		}
	}
}

func (h *HttpRequest) genHttpRequest(requestMethod, url, requestBody string) (*http.Request, error) {
	jsonBody := []byte(requestBody)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequestWithContext(h.ctx, requestMethod, url, bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "create http request failed")
	}

	h.genHttpHeaders(req)

	return req, nil
}

func (h *HttpRequest) requestByForm(method string, params url.Values) ([]byte, error) {
	body := bytes.NewBufferString(params.Encode())

	req, err := http.NewRequest(method, h.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	if len(h.headers) > 0 {
		for k, v := range h.headers {
			req.Header.Add(k, v)
		}
	}

	// Fetch Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func WithHeaders(headers map[string]string) HttpRequestOption {
	return func(h *HttpRequest) {
		h.headers = headers
	}
}

func WithTimeout(timeout time.Duration) HttpRequestOption {
	return func(h *HttpRequest) {
		h.timeout = timeout
	}
}

func WithRequestBody(body string) HttpRequestOption {
	return func(h *HttpRequest) {
		h.requestBody = body
	}
}

func WithProxy(proxy string) HttpRequestOption {
	return func(h *HttpRequest) {
		h.proxy = proxy
	}
}

func fixHost(host string) string {
	if len(host) == 0 {
		return ""
	}

	if host[len(host)-1:] == "/" {
		host = host[:len(host)-1]
	}
	return host
}

func fixRelativeUrl(relativeUrl string) string {
	if len(relativeUrl) == 0 {
		return ""
	}

	if relativeUrl[0:1] != "/" {
		relativeUrl = "/" + relativeUrl
	}
	return relativeUrl
}
