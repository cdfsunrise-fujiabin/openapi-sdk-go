package httpClient

import (
	"context"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type QueryStockRecordResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any `json:"data"`
}
/*QueryStockRecord
 *Description: 查询库存操作记录
 * @param: channelId string channelId 必填项
 * @param: requestNo string requestNo 必填项
 * @return: *QueryStockRecordResponse
*/
func (t *CdfSunriseRequestClient) QueryStockRecord(ctx context.Context, channelId string, requestNo string) (*QueryStockRecordResponse, error) {
	headers := GenHeaders(nil)
	
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/query/stock/record?channelId=%v&requestNo=%v", channelId, requestNo), exHttp.WithHeaders(headers)).GetUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity QueryStockRecordResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
