package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type ZltOrderCreatedResponse struct {
	Success bool `json:"success"`
}

type ZltOrderCreatedData struct {
	JsonData  string `json:"jsonData"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
}

/*V1ZltOrderCreated
 *Description: 中旅通订单推送
 * @param: body ZltOrderCreatedData 中旅通订单推送数据结构,需要到handler里处理, timestamp & sign 都在header里, requestBody 从requestbody里取 必填项
 * @return: *ZltOrderCreatedResponse
 */
func (t *CdfSunriseRequestClient) V1ZltOrderCreated(ctx context.Context, authToken string, body ZltOrderCreatedData) (*ZltOrderCreatedResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/zlt/order/created"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity ZltOrderCreatedResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
