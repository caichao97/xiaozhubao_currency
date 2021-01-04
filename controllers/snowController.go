package controllers

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	helper "xiaozhubao_currency/Helper"
	models "xiaozhubao_currency/Models"
)

type SnowController struct {
	BaseController
	ListMap map[string]string
}

func (snow *SnowController) GetId() {
	workId, _ := snow.GetInt64("workId",0)
	snow.ListMap = make(map[string]string)
	if workId == 0 {
		snow.outJson(snow.ListMap,201,"FAIL workId is must")
	}
	snowWorker, _ := helper.NewWorker(workId) //初始化工作节点
	uuid := snowWorker.GetId()
	workids := models.Workids{Id: uuid,WorkId: workId}
	orms :=orm.NewOrm()
	_,errors :=orms.Insert(&workids)
	if errors != nil {
		snow.outJson(snow.ListMap,201,"FAIL")
	}
	snow.ListMap["ID"] = strconv.FormatInt(uuid,10)
	snow.outJson(snow.ListMap,200,"SUCCESS")
}