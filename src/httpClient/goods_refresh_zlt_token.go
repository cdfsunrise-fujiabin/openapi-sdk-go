package httpClient

import (
	"context"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type RefreshZltTokenResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      string `mapstructure:"data"`
}

/*RefreshZltToken
 *Description: 更新中旅通accessToken
 * @param: appid string appid 必填项
 * @return: *RefreshZltTokenResponse
 */
func (t *CdfSunriseRequestClient) RefreshZltToken(ctx context.Context, authToken string, appid string) (*RefreshZltTokenResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/refresh/zlt/token?appid=%v", appid), exHttp.WithHeaders(headers)).GetUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity RefreshZltTokenResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
