package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1MallInvoiceApplyResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      string `mapstructure:"data"`
}

type BaseRequest struct {
	Appid             string `json:"appid"`
	Data              string `json:"data"`
	DataEncryptMethod string `json:"dataEncryptMethod"`
	Sign              string `json:"sign"`
	SignEncryptMethod string `json:"signEncryptMethod"`
	Timestamp         string `json:"timestamp"`
}

/*V1MallInvoiceApply
 *Description: 【商户入驻】- 订单发票开票申请回执
 * @param: body BaseRequest BaseRequest 必填项
 * @return: *V1MallInvoiceApplyResponse
 */
func (t *CdfSunriseRequestClient) V1MallInvoiceApply(ctx context.Context, authToken string, body BaseRequest) (*V1MallInvoiceApplyResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/mall/invoice/apply"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1MallInvoiceApplyResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
