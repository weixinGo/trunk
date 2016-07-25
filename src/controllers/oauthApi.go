package controllers

import (
	"github.com/astaxie/beego"
)

type OauthApiController struct {
	beego.Controller
}

func (c *OauthApiController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
