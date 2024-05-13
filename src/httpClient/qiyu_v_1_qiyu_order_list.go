package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type QiYuOrderListResponse struct {
	Code	int `json:"code"`
	Desc	string `json:"desc"`
	RequestId	string `json:"requestId"`
	Result	
}


type QiYuOrderListQuery struct {
	Params	object `json:"params"`
	PlgCode	string `json:"plgCode"`
	Token	string `json:"token"`
	UserId	string `json:"userId"`
}

/*V1QiyuOrderList
 *Description: 订单列表查询
 * @param: body QiYuOrderListQuery QiYuOrderListQuery 必填项
 * @return: *QiYuOrderListResponse
*/
func (t *CdfSunriseRequestClient) V1QiyuOrderList(ctx context.Context, body QiYuOrderListQuery) (*QiYuOrderListResponse, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/qiyu/order/list"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity QiYuOrderListResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
