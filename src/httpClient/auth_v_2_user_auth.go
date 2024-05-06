package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V2UserAuthResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any `json:"data"`
}

type OpenAuthReq struct {
	Appid	string `json:"appid"`
	Password	string `json:"password"`
}

/*V2UserAuth
 *Description: 
 * @param: body OpenAuthReq OpenAuthReq 必填项
 * @return: *V2UserAuthResponse
*/
func (t *CdfSunriseRequestClient) V2UserAuth(ctx context.Context, body OpenAuthReq) (*V2UserAuthResponse, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v2/user/auth"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity V2UserAuthResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
