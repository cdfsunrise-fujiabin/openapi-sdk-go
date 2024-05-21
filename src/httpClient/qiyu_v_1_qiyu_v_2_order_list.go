package httpClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cdfsunrise-fujiabin/openapi-sdk-go/src/utils/exHttp"
	"github.com/mitchellh/mapstructure"
)

type QiYuOrderApiResult struct {
	Data	QiYuOrderApiData
}

type QiYuOrderApiData struct {
	Banner	string `json:"banner"`
	CurrentPage	int `json:"current_page"`
	EmptyListHint	string `json:"empty_list_hint"`
	Introducer	string `json:"introducer"`
	PageSize	int `json:"page_size"`
	Tab1	QiYuOrderApiTab
}

	Tab2	QiYuOrderApiTab
}

	Tab3	QiYuOrderApiTab
}

	Tab4	QiYuOrderApiTab
}

	TabId	int `json:"tab_id"`


type QiYuOrderListQuery struct {
	Params	any `json:"params"`
	PlgCode	string `json:"plgCode"`
	Token	string `json:"token"`
	UserId	string `json:"userId"`
}

/*V1QiyuV2OrderList
 *Description: 订单列表查询
 * @param: body QiYuOrderListQuery QiYuOrderListQuery 必填项
 * @return: *QiYuOrderApiResult
*/
func (t *CdfSunriseRequestClient) V1QiyuV2OrderList(ctx context.Context, body QiYuOrderListQuery) (*QiYuOrderApiResult, error) {
	headers := GenHeaders(nil)
	
	marshal, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    respMap, err := exHttp.NewHttpRequest(ctx, t.host, fmt.Sprintf("/v1/qiyu/v2/order/list"), exHttp.WithHeaders(headers), exHttp.WithRequestBody(string(marshal))).PostUnmarshal()
	
    if err != nil {
        return nil, err
    }

	var respEntity QiYuOrderApiResult
	err = mapstructure.Decode(respMap, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
