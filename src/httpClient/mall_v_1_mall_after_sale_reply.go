package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1MallAfterSaleReplyResponse struct {
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

/*V1MallAfterSaleReply
 *Description: 【商户入驻】- 售后信息回执
 * @param: body BaseRequest BaseRequest 必填项
 * @return: *V1MallAfterSaleReplyResponse
 */
func (t *CdfSunriseRequestClient) V1MallAfterSaleReply(ctx context.Context, authToken string, body BaseRequest) (*V1MallAfterSaleReplyResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/mall/after/sale/reply"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1MallAfterSaleReplyResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
