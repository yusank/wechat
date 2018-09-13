// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/changxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

package statistics

import (
	"github.com/changxuehong/wechat/mp"
	"github.com/changxuehong/wechat/mp/shakearound/device"
)

// 以设备为维度的数据统计接口
func Device(clt *mp.Client, deviceIdentifier *device.DeviceIdentifier, beginDate, endDate int64) (data []StatisticsBase, err error) {
	request := struct {
		DeviceIdentifier *device.DeviceIdentifier `json:"device_identifier,omitempty"`
		BeginDate        int64                    `json:"begin_date"`
		EndDate          int64                    `json:"end_date"`
	}{
		DeviceIdentifier: deviceIdentifier,
		BeginDate:        beginDate,
		EndDate:          endDate,
	}

	var result struct {
		mp.Error
		Data []StatisticsBase `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/device?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	data = result.Data
	return
}
