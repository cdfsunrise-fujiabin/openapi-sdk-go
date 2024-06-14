package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type SyncGoodPriceResponse struct {
	RequestId string     `mapstructure:"requestId"`
	Code      int        `mapstructure:"code"`
	Message   string     `mapstructure:"message"`
	Data      []GoodResp `mapstructure:"data"`
}

type SyncGoodPriceReq struct {
	GoodPriceList []GoodPriceInfo `json:"goodPriceList"`
	Supplier      string          `json:"supplier"`
}

type GoodPriceInfo struct {
	LefoxId      string `json:"lefoxId"`
	MerchantId   string `json:"merchantID"`
	MerchantName string `json:"merchantName"`
	SalePrice    string `json:"salePrice"`
}

type GoodResp struct {
	ErrInfo string `json:"errInfo"`
	LefoxId string `json:"lefoxId"`
	Success bool   `json:"success"`
}

/*SyncGoodPrice
 *Description: 同步商品价格信息
 * @param: body SyncGoodPriceReq SyncGoodPriceReq 必填项
 * @return: *SyncGoodPriceResponse
 */
func (t *CdfSunriseRequestClient) SyncGoodPrice(ctx context.Context, authToken string, body SyncGoodPriceReq) (*SyncGoodPriceResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

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
