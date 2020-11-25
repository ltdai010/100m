package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"push-data/models"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid, err := models.AddOne(ob)
	if err != nil {
		o.Data["json"] = err
		o.ServeJSON()
		return
	}
	o.Data["json"] = map[string]string{"ObjectId": fmt.Sprint(objectid)}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}


// @Title GetPaginateObjectWithName
// @Description find object by objectid
// @Param	objectName		query 	string		true		"the object name you want to get"
// @Param	page		query	int		true		"the page"
// @Param	count		query	int		true		"the page length"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /page-object-names [get]
func (o *ObjectController) GetPaginateObjectWithName() {
	objectName := o.GetString("objectName")
	page, _ := o.GetInt32("page")
	count, _ := o.GetInt32("count")
	ob, err := models.GetPaginateObjectWithObjectName(objectName, page, count)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = ob
	}
	o.ServeJSON()
}