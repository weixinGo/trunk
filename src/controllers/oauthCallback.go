package controllers

import (
	"github.com/astaxie/beego"
	"../wxHelper"
	"fmt"
)

type OauthCallbackController struct {
	beego.Controller
}

func (c *OauthCallbackController) Get() {
	code := c.GetString("code")
	state := c.GetString("state")
	fmt.Println("code :", code , "; state : ", state)
	openId, info, err := wxHelper.GetUserInfo(code)
	if err != nil{
		fmt.Println("error : " , err)
	}
	fmt.Println("openId : " , openId)
	if info != nil {
		fmt.Println("info : " , *info)
	}
	c.Redirect("/static/html/test_oauthEnd.html", 302)
}
