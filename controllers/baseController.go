package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type JSONS struct {
	Code int
	Msg  string
	List map[string]string
}

func (base BaseController) outJson(list map[string]string,code int,msg string){
	base.Data["json"] = &JSONS{code,msg,list}
	base.ServeJSON()
}
