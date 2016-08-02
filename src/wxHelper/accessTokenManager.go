package WxHelper

import(
    "time"
)

tokens :=make(map[string]*AccessToken)

func GetToken(openId string, secret string)(*AccessToken, err){
    token := tokens[openId]
    if token == nil || token.IsInvalid(){
        t, err := GetTokenFromWX(openId, secret)
        if err != nil {
            return nil,err
        }
        tokens[openId] = t
        return t
    }
    return token
}

func UpdateToken(openId string, accessToken string, expiresIn int64){
    token := new(accessToken)
    token.Access_token = accessToken
    token.Expires_in = expiresIn
    token.UpdateTime = time.Now().Unix()
    tokens[openId] = token
}