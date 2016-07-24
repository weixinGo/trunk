package main

import (
//	_ "./routers"
//	"github.com/astaxie/beego"
	"fmt"
	"./wxHelper"
)

func main() {
	//beego.Run()
	wxHelper.SetOpenId("testOpenid")
	wxHelper.SetSecret("testSecret")
	oUrl,_ := wxHelper.WXOauth2(false, "http://wangqiao.gaiay.net.cn/wxServer?aaa=斯蒂芬撒")
	fmt.Println("oUrl:", oUrl)
}

