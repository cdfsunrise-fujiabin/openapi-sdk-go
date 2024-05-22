package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1UserAccountQueryResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

/*V1UserAccountQuery
 *Description: 根据会员凭证查询用户及百花可用总账
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1UserAccountQueryResponse
 */
func (t *CdfSunriseRequestClient) V1UserAccountQuery(ctx context.Context, body OpenDataReq) (*V1UserAccountQueryResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/user/accountQuery"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1UserAccountQueryResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
