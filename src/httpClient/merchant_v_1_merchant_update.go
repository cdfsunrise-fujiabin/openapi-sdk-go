package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1MerchantUpdateResponse struct {
	RequestId string `mapstructure:"requestId"`
	Code      int    `mapstructure:"code"`
	Message   string `mapstructure:"message"`
	Data      any    `mapstructure:"data"`
}

type MerchantUpdateReq struct {
	Appid               string `json:"appid"`
	ChannelId           string `json:"channelId"`
	DistributionChannel string `json:"distributionChannel"`
	ExpireTime          int    `json:"expireTime"`
	ExtInfo             string `json:"extInfo"`
	GoodsChannelId      string `json:"goodsChannelId"`
	MerchantName        string `json:"merchantName"`
	OrderSource         string `json:"orderSource"`
	Password            string `json:"password"`
	PostHost            string `json:"postHost"`
	PostPath            string `json:"postPath"`
}

/*V1MerchantUpdate
 *Description: 商户信息编辑
 * @param: body MerchantUpdateReq MerchantUpdateReq 必填项
 * @return: *V1MerchantUpdateResponse
 */
func (t *CdfSunriseRequestClient) V1MerchantUpdate(ctx context.Context, authToken string, body MerchantUpdateReq) (*V1MerchantUpdateResponse, error) {
	headers := GenHeaders(map[string]string{
		"Authorization": authToken,
	})

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/merchant/update"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()

	if err != nil {
		return nil, err
	}

	var respEntity V1MerchantUpdateResponse
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
