package httpClient

import (
	"context"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

/*DelMerchantCache
 *Description: 删除缓存
 * @param: appid string  必填项
 * @return: *Response
 */
func (t *CdfSunriseRequestClient) DelMerchantCache(ctx context.Context, appid string) (*Response, error) {
	headers := GenHeaders(nil)

	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/del/merchant/cache?appid=%v", appid), exHttp.WithHeaders(headers)).GetUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity Response
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
