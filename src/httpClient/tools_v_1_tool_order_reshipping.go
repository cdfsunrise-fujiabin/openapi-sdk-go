package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1ToolOrderReshippingResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

type ToolReShipmentReq struct {
	OrderNo []string `json:"orderNo"`
}

/*V1ToolOrderReshipping
 *Description: 订单发货重推到商户
 * @param: body ToolReShipmentReq ToolReShipmentReq 必填项
 * @return: *V1ToolOrderReshippingResponse
 */
func (t *CdfSunriseRequestClient) V1ToolOrderReshipping(ctx context.Context, authToken string, body ToolReShipmentReq) (*V1ToolOrderReshippingResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/tool/order/reshipping"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1ToolOrderReshippingResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
