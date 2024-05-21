package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type OpenapiactionResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

type YZFOrderNotifyRequest struct {
	AppId     string `json:"appId"`
	Data      string `json:"data"`
	Method    string `json:"method"`
	RequestNo string `json:"requestNo"`
	Timestamp string `json:"timestamp"`
}

/*Openapiaction
 *Description: 翼支付消息推送
 * @param: action string  必填项
 * @param: body YZFOrderNotifyRequest YZFOrderNotifyRequest 必填项
 * @return: *OpenapiactionResponse
 */
func (t *CdfSunriseRequestClient) Openapiaction(ctx context.Context, action string, body YZFOrderNotifyRequest) (*OpenapiactionResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/openapi/%v", action), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity OpenapiactionResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
