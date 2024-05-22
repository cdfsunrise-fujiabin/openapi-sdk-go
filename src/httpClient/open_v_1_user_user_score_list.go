package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1UserUserScoreListResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

/*V1UserUserScoreList
 *Description: 积分收/支记录
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1UserUserScoreListResponse
 */
func (t *CdfSunriseRequestClient) V1UserUserScoreList(ctx context.Context, body OpenDataReq) (*V1UserUserScoreListResponse, error) {
	headers := GenHeaders(nil)

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/user/userScoreList"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1UserUserScoreListResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
