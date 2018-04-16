package controllers

import (
	"beego_test/ihome/models"
	"github.com/astaxie/beego"

	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
// /api/v1.0/users [post]
func (c *UserController) Reg() {
	//1

	resp := make(map[string]interface{})
	/*
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
*/
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)
//1.得到客户端请求的json数据 post数据
	regRequestMap := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &regRequestMap)
	beego.Info("mobile = ",regRequestMap["mobile"],"passsword",regRequestMap["password"])
//2.判断数据合法性
	if regRequestMap["mobile"] == "" || regRequestMap["password"] == ""||regRequestMap["sms_code"] == ""{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
//3.将数据存入到Mysql数据库 user表中
	user := models.User{}
	user.Mobile = regRequestMap["mobile"].(string)
	user.Password_hash = regRequestMap["password"].(string)
	user.Name = regRequestMap["mobile"].(string)

	o := orm.NewOrm()
	id,err := o.Insert(&user)
	if err != nil{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		beego.Info("insert fail,id = ",id)
		return
	}
	c.SetSession("name",user.Mobile)
	c.SetSession("user_id",id)
	c.SetSession("mobile",user.Mobile)


}

func (c *UserController) Login() {

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.RetData(resp)
	//1.得到客户端请求的json数据 post数据
	regRequestMap := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &regRequestMap)
	beego.Info("mobile = ",regRequestMap["mobile"],"passsword",regRequestMap["password"])
	//2.判断数据合法性
	if regRequestMap["mobile"] == "" || regRequestMap["password"] == ""{
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}
	//3.查询数据
	var user models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Filter("mobile",regRequestMap["mobile"]).One(&user)
	if user.Password_hash != regRequestMap["password"]{
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}


	c.SetSession("name",user.Mobile)
	c.SetSession("user_id",user.Id)
	c.SetSession("mobile",user.Mobile)


}
