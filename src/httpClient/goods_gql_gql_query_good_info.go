package httpClient

import (
	"context"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type GqlQueryGoodInfoResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any `json:"data"`
}
/*GqlQueryGoodInfo
 *Description: 查询商品信息工具
 * @param: fields string 期望返回的字段 必填项
 * @param: channelId string 渠道id 必填项
 * @param: lefoxId string 商品lefoxid 必填项
 * @return: *GqlQueryGoodInfoResponse
*/
func (t *CdfSunriseRequestClient) GqlQueryGoodInfo(ctx context.Context, fields string, channelId string, lefoxId string) (*GqlQueryGoodInfoResponse, error) {
	headers := GenHeaders(nil)
	
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/gql/query/good/info?fields=%v&channelId=%v&lefoxId=%v", fields, channelId, lefoxId), exHttp.WithHeaders(headers)).GetUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity GqlQueryGoodInfoResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
