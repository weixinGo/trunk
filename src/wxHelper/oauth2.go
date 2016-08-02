package wxHelper 

import (
	//str "strings"
	"io/ioutil"
	"fmt"
	"net/url"
	"net/http"
	"errors"
	"strings"
	 "encoding/json"
)


const WX_OAUTH2 = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s#wechat_redirect"
const WX_OUTOKEN="https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
const WX_USERINFO="https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN" 

type OauthToken struct{
	Token string
	OpenId string
	Scope string
}

type WXUserInfo struct{
	Openid string
	Nickname string
	Sex int
	Province string
	City string
	Country string
	Headimgurl string
	//Privilege string
	Unionid string
}
/*
微信授权入口
输入isBase		是否为隐式授权，true为隐式授权，false为显示授权
输入redirectUrl	回调地址
返回 string	微信授权跳转地址
返回 error	异常信息
*/
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
	fmt.Println("redirectUrl:",redirectUrl ,"uo.Encode():", uo.Encode())
	oUrl := fmt.Sprintf(WX_OAUTH2, base_openId, uo.Encode(), scope)
	return oUrl,nil
}

/*
获取用户信息
输入code	微信返回code码
返回 openId	用户openid
返回 *WXUserInfo 用户详细信息, 隐式授权时，返回nil
返回 error	异常信息
*/
func GetUserInfo(code string) (string, *WXUserInfo,error){
	if len(code) < 1{
		return "", nil, errors.New("返回code异常")
	}
	token, err := getToken(code)
	if err != nil{
		return "",nil,err
	}
	// 隐式授权直接返回openId
	if strings.Compare(token.Scope, "snsapi_base")==0{
		return token.OpenId, nil, nil
	}
	// 如果是显示授权，则获取详细信息
	var info *WXUserInfo
	info, err = getwxUserInfo(token.Token)
	return token.OpenId, info, err 
}

/*
获取微信授权token
输入code	微信返回code码
返回 *OauthToken	获取微信授权access_token返回的基础信息，包括token、openid、scope
返回 error	异常信息
*/
func getToken(code string) (*OauthToken, error){
	url := fmt.Sprintf(WX_OUTOKEN, base_openId, base_secret, code)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return parseToken(string(body))	
}

/*
解析获取微信授权token的返回值
输入respons 微信返回值
返回 *OauthToken	获取微信授权access_token返回的基础信息，包括token、openid、scope
返回 error	异常信息
*/
func parseToken(respons string) (*OauthToken, error){
	err := checkResponse("调用微信获取token接口异常", respons)
	if err != nil{
		return nil, err
	}
	dec := json.NewDecoder(strings.NewReader(respons))
	var dts map[string]interface{} 
	err = dec.Decode(&dts)
	if err != nil {
    	return nil, err
	}

	token := new (OauthToken)
	token.Token = dts["access_token"].(string)
	token.OpenId = dts["openid"].(string)
	token.Scope = dts["scope"].(string) 
	
	return token, nil 
}

func getwxUserInfo(token string)(*WXUserInfo, error){
	url := fmt.Sprintf(WX_USERINFO, token, base_openId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return paresUserInfo(string(body))
}

func paresUserInfo(response string)(*WXUserInfo, error){
	err := checkResponse("调用微信接口获取用户信息异常", response)
	if err != nil{
		return nil, err
	}
	info := new(WXUserInfo)
	dec := json.NewDecoder(strings.NewReader(response))
	err = dec.Decode(info)
	if err != nil {
    	return nil, err
	}
	return info,nil
}

