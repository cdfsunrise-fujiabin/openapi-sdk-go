package httpClient

import (
	"context"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type QueryStockRecordResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

/*QueryStockRecord
 *Description: 查询库存操作记录
 * @param: channelId string channelId 必填项
 * @param: requestNo string requestNo 必填项
 * @return: *QueryStockRecordResponse
 */
func (t *CdfSunriseRequestClient) QueryStockRecord(ctx context.Context, authToken string, channelId string, requestNo string) (*QueryStockRecordResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

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
