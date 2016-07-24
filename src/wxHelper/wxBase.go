package wxHelper 

import (

)

var openId string
var secret string

func SetOpenId(oi string){
	openId = oi
}

func SetSecret(sc string){
	secret = sc
}

func isInited() bool{
	if len(openId) <1 || len(secret) < 1{
		return false
	}
	return true
}

