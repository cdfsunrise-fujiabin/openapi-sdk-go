package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type SyncGoodPriceResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      []GoodResp `json:"data"`
}

type SyncGoodPriceReq struct {
	Supplier	string `json:"supplier"`
	GoodPriceList	[]GoodPriceInfo `json:"goodPriceList"`
}

/*SyncGoodPrice
 *Description: 同步商品价格信息
 * @param: body SyncGoodPriceReq SyncGoodPriceReq 必填项
 * @return: *SyncGoodPriceResponse
*/
func (t *CdfSunriseRequestClient) SyncGoodPrice(ctx context.Context, body SyncGoodPriceReq) (*SyncGoodPriceResponse, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/sync/good/price"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity SyncGoodPriceResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}