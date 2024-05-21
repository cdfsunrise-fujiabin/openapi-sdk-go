package httpClient

import (
	"context"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type QiYuGetTokenApiResponse struct {
	Expires int    `json:"expires"`
	Rlt     int    `json:"rlt"`
	Token   string `json:"token"`
}

/*V1QiyuGetToken
 *Description: 获取token
 * @param: appid string appid 必填项
 * @param: appsecret string appsecret 必填项
 * @return: *QiYuGetTokenApiResponse
 */
func (t *CdfSunriseRequestClient) V1QiyuGetToken(ctx context.Context, appid string, appsecret string) (*QiYuGetTokenApiResponse, error) {
	headers := GenHeaders(nil)

	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/qiyu/get_token?appid=%v&appsecret=%v", appid, appsecret), exHttp.WithHeaders(headers)).GetUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity QiYuGetTokenApiResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
