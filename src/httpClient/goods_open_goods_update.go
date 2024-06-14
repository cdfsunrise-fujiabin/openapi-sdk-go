package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type OpenGoodsUpdateResponse struct {
	RequestId string     `mapstructure:"requestId"`
	Code      int        `mapstructure:"code"`
	Message   string     `mapstructure:"message"`
	Data      []GoodResp `mapstructure:"data"`
}

type OpenGoodUpdateReq struct {
	BuyType      int          `json:"buyType"`
	ChannelId    string       `json:"channelId"`
	GoodListInfo []UpdateGood `json:"goodListInfo"`
}

type UpdateGood struct {
	LefoxId   string `json:"lefoxId"`
	SellState int    `json:"sellState"`
	Type      string `json:"type"`
}

/*OpenGoodsUpdate
 *Description: 开放平台商品信息通知
 * @param: body OpenGoodUpdateReq OpenGoodUpdateReq 必填项
 * @return: *OpenGoodsUpdateResponse
 */
func (t *CdfSunriseRequestClient) OpenGoodsUpdate(ctx context.Context, authToken string, body OpenGoodUpdateReq) (*OpenGoodsUpdateResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

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
