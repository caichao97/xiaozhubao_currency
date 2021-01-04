package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"xiaozhubao_currency/controllers"
)

func init() {
    beego.Router("/getId", &controllers.SnowController{},"*:GetId")
}
