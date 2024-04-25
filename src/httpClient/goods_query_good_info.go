package httpClient

import (
	"cdfsunrise.com/openapi/src/utils/exHttp"
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type QueryGoodInfoResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      string `json:"data"`
}

/*QueryGoodInfo
 *Description: 查询商品信息工具
 * @param: channelId string 渠道id 必填项
 * @param: lefoxId string 商品lefoxid 必填项
 * @return: *QueryGoodInfoResponse
*/
func (t *CdfSunriseRequestClient) QueryGoodInfo(ctx context.Context, authToken string, channelId string, lefoxId string) (*QueryGoodInfoResponse, error) {
	headers := GenHeaders(map[string]string{
		"token": authToken,
	})
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/query/good/info?channelId=%v&lefoxId=%v", channelId, lefoxId), exHttp.WithHeaders(headers)).GetUnmarshal()
	if err != nil {
		return nil, err
	}

	var respEntity QueryGoodInfoResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}