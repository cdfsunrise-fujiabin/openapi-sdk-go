package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type SyncGoodStockResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

type SyncGoodsStockReq struct {
	ChannelId        string           `json:"channelId"`
	GoodsStockList   []GoodsStockItem `json:"goodsStockList"`
	RequestMessageId string           `json:"requestMessageId"`
}

type GoodsStockItem struct {
	GoodsId string `json:"goodsId"`
	LefoxId string `json:"lefoxId"`
	Num     int    `json:"num"`
	Stock   int    `json:"stock"`
	Type    int    `json:"type"`
}

/*SyncGoodStock
 *Description: 同步库存信息
 * @param: body SyncGoodsStockReq SyncGoodsStockReq 必填项
 * @return: *SyncGoodStockResponse
 */
func (t *CdfSunriseRequestClient) SyncGoodStock(ctx context.Context, authToken string, body SyncGoodsStockReq) (*SyncGoodStockResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/sync/good/stock"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity SyncGoodStockResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
