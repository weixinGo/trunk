package main

import (
	_ "./src/routers"
	"github.com/astaxie/beego"
	"./src/wxHelper"
)

func main() {
	wxHelper.SetOpenId("wx9c3ab6e8ce8f1e8b")
	wxHelper.SetSecret("d4ed5e5977b289434b3274a96fd593cd")
	beego.Run()
}

