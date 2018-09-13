// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/changxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package promotion

import (
	"github.com/changxuehong/wechat/mch"
)

// 查询代金券信息.
func QueryCoupon(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/promotion/query_coupon", req)
}
