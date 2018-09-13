// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/changxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mass

import (
	"net/http"

	"github.com/changxuehong/wechat/mp"
)

type Client mp.Client

func NewClient(srv mp.AccessTokenServer, clt *http.Client) *Client {
	return (*Client)(mp.NewClient(srv, clt))
}
