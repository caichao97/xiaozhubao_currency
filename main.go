package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "xiaozhubao_currency/routers"
)

func main() {
	beego.Run()
}

