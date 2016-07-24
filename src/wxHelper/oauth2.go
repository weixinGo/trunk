package wxHelper 

import (
	//str "strings"
	"io/ioutil"
	"fmt"
	"net/url"
	"net/http"
	"errors"
)


const WX_OAUTH2 = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s#wechat_redirect"
const WX_OUTOKEN="https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"

type WXUserInfo struct{
	Openid string
	Nickname string
	Sex string
	Province string
	City string
	Country string
	Headimgurl string
	Privilege string
	Unionid string
}

func WXOauth2(isBase bool, redirectUrl string) (string, error){
	if !isInited(){
		return "", errors.New("系统没有初始化") 
	}
	if len(redirectUrl) <1 {
		return "", errors.New("回调地址为空")
	}
	uo, err := url.ParseQuery(redirectUrl)
	if err != nil{
		return "", errors.New("回调地址不合法") 
	}
	scope := "snsapi_userinfo" 
	if isBase{
		scope = "snsapi_base"
	}
	oUrl := fmt.Sprintf(WX_OAUTH2, openId, uo.Encode(), scope)
	return oUrl,nil
}

func GetUserInfo(code string) (bool, *WXUserInfo,error){
	if len(code) < 1{
		return false, nil, errors.New("返回code异常")
	}
	tokenRsp, err := getToken(code)
	if err != nil{
		return false,nil,err
	}
	// 获取token和code
	parseToken(tokenRsp)
	// 如果是显示授权，则获取详细信息
	
	return false,nil,errors.New("未实现")
}

func parseToken(respons string){

}

func getToken(code string) (string, error){
	url := fmt.Sprintf(WX_OUTOKEN, openId, secret, code)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil	
}

