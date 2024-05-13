package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type SyncCdfStockResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any `json:"data"`
}

type SyncCdfGoodsStockReq struct {
	Items	[]SyncCdfGoodsStockItem `json:"items"`
	OccurTime	string `json:"occurTime"`
}

type SyncCdfGoodsStockItem struct {
	BeforeQuantity	int `json:"beforeQuantity"`
	BuyTypeAgg	int `json:"buyTypeAgg"`
	Quantity	int `json:"quantity"`
	RTmerchantId	string `json:"rTMerchantId"`
	RTskuId	string `json:"rTSkuId"`
	StockChannelId	string `json:"stockChannelId"`
}
/*SyncCdfStock
 *Description: 同步cdf中免会员商品库存
 * @param: body SyncCdfGoodsStockReq SyncCdfGoodsStockReq 必填项
 * @return: *SyncCdfStockResponse
*/
func (t *CdfSunriseRequestClient) SyncCdfStock(ctx context.Context, body SyncCdfGoodsStockReq) (*SyncCdfStockResponse, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/sync/cdf/stock"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity SyncCdfStockResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
