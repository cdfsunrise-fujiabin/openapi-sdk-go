package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1CdfGoodsResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

/*V1CdfGoods
 *Description: 获取cdf商品列表
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1CdfGoodsResponse
 */
func (t *CdfSunriseRequestClient) V1CdfGoods(ctx context.Context, authToken string, body OpenDataReq) (*V1CdfGoodsResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/cdf/goods"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1CdfGoodsResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
