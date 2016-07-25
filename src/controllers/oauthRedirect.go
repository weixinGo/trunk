package controllers

import (
	"github.com/astaxie/beego"
	"../wxHelper"
	"fmt"
)

type OauthRedirectController struct {
	beego.Controller
}

func (c *OauthRedirectController) Get() {
	isBase,_ := c.GetBool("isBase", true)
	url, err := wxHelper.WXOauth2(isBase, "http://wangqiao.gaiay.net.cn/weixin/oauth/callback?1=1")
	if err != nil {
		fmt.Println("error:", err)		
	}
	fmt.Println("redirect url:", url)
	c.Redirect(url, 302)
	return
}
