package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1StockCountResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      []OpenGoodsStock `json:"data"`
}

type OpenDataReq struct {
	Appid	string `json:"appid"`
	Data	string `json:"data"`
	DataEncryptMethod	string `json:"dataEncryptMethod"`
	Sign	string `json:"sign"`
	SignEncryptMethod	string `json:"signEncryptMethod"`
	Timestamp	string `json:"timestamp"`
}


type OpenGoodsStock struct {
	LefoxId	string `json:"lefoxId"`
	Stock	int `json:"stock"`
}

/*V1StockCount
 *Description: 开放平台批量查询商品库存
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1StockCountResponse
*/
func (t *CdfSunriseRequestClient) V1StockCount(ctx context.Context, body OpenDataReq) (*V1StockCountResponse, error) {
	headers := GenHeaders(nil)
	
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
