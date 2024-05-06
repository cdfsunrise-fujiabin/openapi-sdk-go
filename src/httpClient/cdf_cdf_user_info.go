package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type CdfUserInfoResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any `json:"data"`
}

type CdfUserInfoReq struct {
	ChannelId	string `json:"channelId"`
	Token	string `json:"token"`
	UserId	string `json:"userId"`
}

/*CdfUserInfo
 *Description: sso获取用户信息
 * @param: body CdfUserInfoReq CdfUserInfoReq 必填项
 * @return: *CdfUserInfoResponse
*/
func (t *CdfSunriseRequestClient) CdfUserInfo(ctx context.Context, body CdfUserInfoReq) (*CdfUserInfoResponse, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/cdf/userInfo"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity CdfUserInfoResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
