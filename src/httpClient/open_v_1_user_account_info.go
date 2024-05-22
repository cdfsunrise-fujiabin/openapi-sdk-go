package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1UserAccountInfoResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

type OpenDataReq struct {
	Appid             string `json:"appid"`
	Data              string `json:"data"`
	DataEncryptMethod string `json:"dataEncryptMethod"`
	Sign              string `json:"sign"`
	SignEncryptMethod string `json:"signEncryptMethod"`
	Timestamp         string `json:"timestamp"`
}

/*V1UserAccountInfo
 *Description: 解密qrCode获取用户信息
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1UserAccountInfoResponse
 */
func (t *CdfSunriseRequestClient) V1UserAccountInfo(ctx context.Context, body OpenDataReq) (*V1UserAccountInfoResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/user/accountInfo"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1UserAccountInfoResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
