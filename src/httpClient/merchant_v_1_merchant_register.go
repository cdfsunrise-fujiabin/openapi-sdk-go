package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1MerchantRegisterResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

type MerchantRegisterReq struct {
	BusinessType int    `json:"businessType"`
	MerchantName string `json:"merchantName"`
}

/*V1MerchantRegister
 *Description: 商户注册
 * @param: body MerchantRegisterReq MerchantRegisterReq 必填项
 * @return: *V1MerchantRegisterResponse
 */
func (t *CdfSunriseRequestClient) V1MerchantRegister(ctx context.Context, body MerchantRegisterReq) (*V1MerchantRegisterResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/merchant/register"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1MerchantRegisterResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
