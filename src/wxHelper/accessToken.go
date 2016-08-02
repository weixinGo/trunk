package wxHelper

import(
    "time"
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type AccessToken struct{
    Access_token string
    Expires_in int64
    UpdateTime int64
}

const WX_ACCESS_TOKEN="https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

func (token *AccessToken)IsTokenIvalid() bool{
    if len(Access_token)>0 && time.Now().Unix() - token.UpdateTime < (token.Expires_in-30)*1000{
        return false
    }
    return true
}

func GetTokenFromWX(openId string, secret string)(*AccessToken, error){
    url := fmt.Sprintf(WX_ACCESS_TOKEN, openId, secret) 
    response, err := http.Get(url)
    defer resp.Body.Close()
    if err != nil {
        return nil,err
    }
    body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
    return parseToken(body)
}


func paresToken(rsp string)(*AccessToken ,error){
    err := checkResponse("调用微信接口获取用户信息异常", rsp)
    if err != nil {
        return nil,err
    }
    dec := json.NewDecoder(strings.NewReader(response))
    token := new(AccessToken)
    token.UpdateTime = time.Now().Unix()   
    err := dec.Decode(token)
    if err != nil {
        return nil, err
    }
    return token, nil
}