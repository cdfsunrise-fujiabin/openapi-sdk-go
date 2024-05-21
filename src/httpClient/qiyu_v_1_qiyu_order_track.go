package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type QiYuOrderTrackApiResult struct {
	Data	QiYuOrderTrackData
}

type QiYuOrderTrackData struct {
	Banner	string `json:"banner"`
	EmptyListHint	string `json:"empty_list_hint"`
	Footer	string `json:"footer"`
	Introducer	string `json:"introducer"`
	ItemList	[]QiYuOrderTrackItem `json:"item_list"`
	Target	string `json:"target"`

type QiYuOrderTrackItem struct {
	Desc	string `json:"desc"`
	IsCurrent	string `json:"is_current"`
	Title	string `json:"title"`


type QiYuOrderTrackQuery struct {
	OrderId	string `json:"orderId"`
}

/*V1QiyuOrderTrack
 *Description: 订单详情物流跟踪
 * @param: body QiYuOrderTrackQuery QiYuOrderTrackQuery 必填项
 * @return: *QiYuOrderTrackApiResult
*/
func (t *CdfSunriseRequestClient) V1QiyuOrderTrack(ctx context.Context, body QiYuOrderTrackQuery) (*QiYuOrderTrackApiResult, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/qiyu/order/track"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity QiYuOrderTrackApiResult
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
