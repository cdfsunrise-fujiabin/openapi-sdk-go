package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1AlipayNotifyResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

type AlipayNotifyRequest struct {
	AppId        string `json:"app_id"`
	BizContent   string `json:"biz_content"`
	Charset      string `json:"charset"`
	MsgMethod    string `json:"msg_method"`
	NotifyId     string `json:"notify_id"`
	Sign         string `json:"sign"`
	UtcTimestamp string `json:"utc_timestamp"`
	Version      string `json:"version"`
}

/*V1AlipayNotify
 *Description: 接收支付宝消息推送
 * @param: body AlipayNotifyRequest AlipayNotifyRequest 必填项
 * @return: *V1AlipayNotifyResponse
 */
func (t *CdfSunriseRequestClient) V1AlipayNotify(ctx context.Context, authToken string, body AlipayNotifyRequest) (*V1AlipayNotifyResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/alipay/notify"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1AlipayNotifyResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
