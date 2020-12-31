package main

import (
	_ "xiaozhubao_currency/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

