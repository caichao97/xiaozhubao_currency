package controllers

import (
	"strconv"
	helper "xiaozhubao_currency/Helper"
)

type SnowController struct {
	BaseController
	ListMap map[string]string
}

func (snow *SnowController) Post(){
	workId, _ := snow.GetInt64("workId",0)
	snow.ListMap = make(map[string]string)
	if workId == 0 {
		snow.outJson(snow.ListMap,201,"FAIL workId is must")
	}
	snowWorker, _ := helper.NewWorker(workId) //初始化工作节点
	snow.ListMap["ID"] = strconv.FormatInt(snowWorker.GetId(),10)
	snow.outJson(snow.ListMap,200,"SUCCESS")
}
