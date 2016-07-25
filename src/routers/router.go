package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/weixin/oauth/", &controllers.OauthRedirectController{})
	beego.Router("/api/weixin/oauth/", &controllers.OauthApiController{})
	beego.Router("/weixin/oauth/callback/", &controllers.OauthCallbackController{})
}
