package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type SyncPriceGoodsResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      string `json:"data"`
}

type SyncPriceGoodsReq struct {
	GoodIds	[]string `json:"goodIds"`
}

/*SyncPriceGoods
 *Description: 需要同步商品价格集合
 * @param: body SyncPriceGoodsReq SyncPriceGoodsReq 必填项
 * @return: *SyncPriceGoodsResponse
*/
func (t *CdfSunriseRequestClient) SyncPriceGoods(ctx context.Context, body SyncPriceGoodsReq) (*SyncPriceGoodsResponse, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/sync/price/goods"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity SyncPriceGoodsResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
