package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1GoodsCategoryResponse struct {
	RequestId string              `mapstructure:"requestId"`
	Code      int                 `mapstructure:"code"`
	Message   string              `mapstructure:"message"`
	Data      []OpenGoodsCategory `mapstructure:"data"`
}

type OpenGoodsCategory struct {
	CategoryInfo CategoryDetail
}

type CategoryDetail struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}

/*V1GoodsCategory
 *Description: 获取分类
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1GoodsCategoryResponse
 */
func (t *CdfSunriseRequestClient) V1GoodsCategory(ctx context.Context, authToken string, body OpenDataReq) (*V1GoodsCategoryResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/goods/category"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1GoodsCategoryResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
