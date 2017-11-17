package controllers

import (
	"GettApi/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Driver
type DriverController struct {
	beego.Controller
}

// @Title Create
// @Description create Driver
// @Param	body		body 	models.Driver	true		"The Driver content"
// @Success 200 {string} models.Driver.Id
// @Failure 403 body is empty
// @router / [post]
func (o *DriverController) Post() {
	var ob models.Driver
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	Driverid := models.AddOneDriver(ob)
	o.Data["json"] = map[string]int{"DriverId": Driverid}
	o.ServeJSON()
}

// @Title Get
// @Description find Driver by Driverid
// @Param	DriverId		path 	string	true		"the Driverid you want to get"
// @Success 200 {Driver} models.Driver
// @Failure 403 :DriverId is empty
// @router /:DriverId [get]
func (o *DriverController) Get() {
	DriverId := o.Ctx.Input.Param(":DriverId")
	if DriverId != "" {
		ob, err := models.GetOneDriver(DriverId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all Drivers
// @Success 200 {Driver} models.Driver
// @Failure 403 :DriverId is empty
// @router / [get]
func (o *DriverController) GetAll() {
	obs := models.GetAllDriver()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the Driver
// @Param	DriverId		path 	string	true		"The Driverid you want to update"
// @Param	body		body 	models.Driver	true		"The body"
// @Success 200 {Driver} models.Driver
// @Failure 403 :DriverId is empty
// @router /:DriverId [put]
func (o *DriverController) Put() {
	//DriverId := o.Ctx.Input.Param(":DriverId")
	var ob models.Driver
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	//S	err := models.UpdateDriver(DriverId, ob.Score)
	// if err != nil {
	// 	o.Data["json"] = err.Error()
	// } else {
	o.Data["json"] = "update success!"
	//	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the Driver
// @Param	DriverId		path 	string	true		"The DriverId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 DriverId is empty
// @router /:DriverId [delete]
func (o *DriverController) Delete() {
	DriverId := o.Ctx.Input.Param(":DriverId")
	models.DeleteDriver(DriverId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
