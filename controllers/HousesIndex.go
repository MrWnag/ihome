package controllers

import (
	"beego_test/ihome/models"
	"github.com/astaxie/beego"
)

type HouseIndexController struct {
	beego.Controller
}
func (this *HouseIndexController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *HouseIndexController) GetHouseIndex() {
	beego.Info("houseIndex resp success")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
}
