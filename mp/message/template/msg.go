// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/changxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package template

import (
	"encoding/json"
)

type TemplateMessage struct {
	ToUser     string `json:"touser"`             // 必须, 接受者OpenID
	TemplateId string `json:"template_id"`        // 必须, 模版ID
	URL        string `json:"url,omitempty"`      // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	TopColor   string `json:"topcolor,omitempty"` // 可选, 整个消息的颜色, 可以不设置

	RawJSONData json.RawMessage `json:"data"` // 必须, JSON 格式的 []byte, 满足特定的模板需求
}

type Template struct {
	TemplateId      string `json:"template_id"`      //模板ID
	Title           string `json:"title"`            //模板标题
	PrimaryIndustry string `json:"primary_industry"` //模板所属行业的一级行业
	DeputyIndustry  string `json:"deputy_industry"`  //模板所属行业的二级行业
	Content         string `json:"content"`          //模板内容
	Example         string `json:"example"`          //模板示例
}

type Industry struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}
