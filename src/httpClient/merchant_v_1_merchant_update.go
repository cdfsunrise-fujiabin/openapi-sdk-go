package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type V1MerchantUpdateResponse struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      any    `json:"data"`
}

type MerchantUpdateReq struct {
	Appid               string `json:"appid"`
	MerchantName        string `json:"merchantName"`
	Password            string `json:"password"`
	GoodsChannelId      string `json:"goodsChannelId"`
	ExpireTime          int    `json:"expireTime"`
	ExtInfo             string `json:"extInfo"`
	PostPath            string `json:"postPath"`
	PostHost            string `json:"postHost"`
	ChannelId           string `json:"channelId"`
	OrderSource         string `json:"orderSource"`
	DistributionChannel string `json:"distributionChannel"`
}

/*V1MerchantUpdate
 *Description: 商户信息编辑
 * @param: body MerchantUpdateReq MerchantUpdateReq 必填项
 * @return: *V1MerchantUpdateResponse
 */
func (t *CdfSunriseRequestClient) V1MerchantUpdate(ctx context.Context, body MerchantUpdateReq) (*V1MerchantUpdateResponse, error) {
	headers := GenHeaders(nil)

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
