package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1ChannelRegisterResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

type ChannelRegisterReq struct {
	ChannelId   string `json:"channelId"`
	ChannelName string `json:"channelName"`
	ChannelType int    `json:"channelType"`
}

/*V1ChannelRegister
 *Description: 渠道库存注册
 * @param: body ChannelRegisterReq ChannelRegisterReq 必填项
 * @return: *V1ChannelRegisterResponse
 */
func (t *CdfSunriseRequestClient) V1ChannelRegister(ctx context.Context, authToken string, body ChannelRegisterReq) (*V1ChannelRegisterResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/channel/register"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1ChannelRegisterResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
