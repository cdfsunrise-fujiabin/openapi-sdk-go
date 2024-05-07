package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type OpenGoodsUpdateResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      []any  `json:"data"`
}

type OpenGoodUpdateReq struct {
	GoodListInfo []any  `json:"goodListInfo"`
	ChannelId    string `json:"channelId"`
	BuyType      int    `json:"buyType"`
}

/*OpenGoodsUpdate
 *Description: 开放平台商品信息通知
 * @param: body OpenGoodUpdateReq OpenGoodUpdateReq 必填项
 * @return: *OpenGoodsUpdateResponse
 */
func (t *CdfSunriseRequestClient) OpenGoodsUpdate(ctx context.Context, body OpenGoodUpdateReq) (*OpenGoodsUpdateResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/open/goods/update"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity OpenGoodsUpdateResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
