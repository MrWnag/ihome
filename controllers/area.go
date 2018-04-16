package controllers

import (
	"beego_test/ihome/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (c *AreaController) GetAreaInfo() {
	//1

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = "OK"
	defer c.RetData(resp)
	o := orm.NewOrm()

	var areas []models.Area

	qs := o.QueryTable("area")
	num, err := qs.All(&areas)

	if err != nil {
		resp["errno"] = 4001
		resp["errmsg"] = "err"
		return
	}

	if num == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "err2"
		return
	}

	resp["errno"] = 0
	resp["errmsg"] = "OK"
	resp["data"] = areas

}
