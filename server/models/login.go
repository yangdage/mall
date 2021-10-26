package models

// 小程序登录凭证校验模型
type AppCode2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// 小程序登录凭证校验,返回的JSON数据包模型
type AppCode2SessionJson struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
