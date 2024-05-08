package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type PingAnOrderResponse struct {
	BizContent string `json:"bizContent"`
	BizSign    string `json:"bizSign"`
	Memo       string `json:"memo"`
	RespCode   string `json:"respCode"`
}

type PingAnOrderQuery struct {
	BizContent string `json:"bizContent"`
	BizSign    string `json:"bizSign"`
	ReqNum     string `json:"reqNum"`
	ReqTime    string `json:"reqTime"`
}

/*PinganOrderQuery
 *Description: 订单状态查询
 * @param: body PingAnOrderQuery PingAnOrderQuery 必填项
 * @return: *PingAnOrderResponse
 */
func (t *CdfSunriseRequestClient) PinganOrderQuery(ctx context.Context, body PingAnOrderQuery) (*PingAnOrderResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/pingan/order/query"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity PingAnOrderResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
