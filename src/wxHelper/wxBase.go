package wxHelper 

import (

)

var base_openId string
var base_secret string

func SetOpenId(oi string){
	base_openId = oi
}

func SetSecret(sc string){
	base_secret = sc
}

func isInited() bool{
	if len(base_openId) <1 || len(base_secret) < 1{
		return false
	}
	return true
}