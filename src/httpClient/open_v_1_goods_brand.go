package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1GoodsBrandResponse struct {
	RequestId string           `json:"requestId"`
	Code      int              `json:"code"`
	Message   string           `json:"message"`
	Data      []OpenGoodsBrand `json:"data"`
}

type OpenGoodsBrand struct {
	BrandInfo BrandDetail
}

type BrandDetail struct {
	Code   string `json:"code"`
	NameCn string `json:"nameCn"`
	NameEn string `json:"nameEn"`
}

/*V1GoodsBrand
 *Description: 获取品牌
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1GoodsBrandResponse
 */
func (t *CdfSunriseRequestClient) V1GoodsBrand(ctx context.Context, body OpenDataReq) (*V1GoodsBrandResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/goods/brand"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1GoodsBrandResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
