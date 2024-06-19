package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1QueryGoodsBrandResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      string `mapstructure:"data"`
}

type OpenDataReq struct {
	Appid             string `json:"appid"`
	Data              string `json:"data"`
	DataEncryptMethod string `json:"dataEncryptMethod"`
	Sign              string `json:"sign"`
	SignEncryptMethod string `json:"signEncryptMethod"`
	Timestamp         string `json:"timestamp"`
}

/*V1QueryGoodsBrand
 *Description: 开放平台商品品牌查询
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1QueryGoodsBrandResponse
 */
func (t *CdfSunriseRequestClient) V1QueryGoodsBrand(ctx context.Context, authToken string, body OpenDataReq) (*V1QueryGoodsBrandResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/query/goodsBrand"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1QueryGoodsBrandResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
