package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1UserCreateAndTagResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

/*V1UserCreateAndTag
 *Description: 创建用户账户并且打上企业标签
 * @param: body OpenDataReq OpenDataReq 必填项
 * @return: *V1UserCreateAndTagResponse
 */
func (t *CdfSunriseRequestClient) V1UserCreateAndTag(ctx context.Context, authToken string, body OpenDataReq) (*V1UserCreateAndTagResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/user/createAndTag"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1UserCreateAndTagResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
