// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/changxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package payutil

import (
	"github.com/changxuehong/wechat/mch"
)

// 测速上报.
func Report(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/payitil/report", req)
}
