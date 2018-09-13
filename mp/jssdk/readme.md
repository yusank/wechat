### 获取 jsapi_ticket 示例
```Go
package main

import (
	"fmt"

	"github.com/changxuehong/wechat/mp"
	"github.com/changxuehong/wechat/mp/jssdk"
)

var AccessTokenServer = mp.NewDefaultAccessTokenServer("appid", "appsecret", nil)
var mpClient = mp.NewClient(AccessTokenServer, nil)
var TicketServer = jssdk.NewDefaultTicketServer(mpClient)

func main() {
	fmt.Println(TicketServer.Ticket())
}
```