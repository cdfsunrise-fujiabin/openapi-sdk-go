package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1QueryGoodsStockResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      string `mapstructure:"data"`
}

/*V1QueryGoodsStock
 *Description: 开放平台查询商品库存
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1QueryGoodsStockResponse
 */
func (t *CdfSunriseRequestClient) V1QueryGoodsStock(ctx context.Context, authToken string, body OpenDataReq) (*V1QueryGoodsStockResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/query/goodsStock"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1QueryGoodsStockResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
