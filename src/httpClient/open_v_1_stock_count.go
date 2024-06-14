package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1StockCountResponse struct {
	RequestId string           `mapstructure:"requestId"`
	Code      int              `mapstructure:"code"`
	Message   string           `mapstructure:"message"`
	Data      []OpenGoodsStock `mapstructure:"data"`
}

type OpenGoodsStock struct {
	LefoxId string `json:"lefoxId"`
	Stock   int    `json:"stock"`
}

/*V1StockCount
 *Description: 开放平台批量查询商品库存
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1StockCountResponse
 */
func (t *CdfSunriseRequestClient) V1StockCount(ctx context.Context, authToken string, body OpenDataReq) (*V1StockCountResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/stock/count"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1StockCountResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
