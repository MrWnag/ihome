package controllers

import (
	"beego_test/ihome/models"
	"github.com/astaxie/beego"

)

type SessionController struct {
	beego.Controller
}
func (this *SessionController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *SessionController) GetSessionData() {
beego.Info("========/api/v1.0/session sucess")

resp := make(map[string]interface{})


defer this.RetData(resp)
nameMap := make(map[string]interface{})

nameMap["name"] = this.GetSession("name").(string)
if nameMap["name"] == nil{
	resp["errno"] = models.RECODE_SESSIONERR
	resp["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
	return
}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = nameMap
}
func (this *SessionController) DelSessionData() {
	beego.Info("========/api/v1.0/session del sucess")

	resp := make(map[string]interface{})


	defer this.RetData(resp)

	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return


}