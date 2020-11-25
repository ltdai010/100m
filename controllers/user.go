package controllers

import (
	"fmt"
	"push-data/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid, err := models.AddUser(user)
	if err != nil {
		u.Data["json"] = err
		u.ServeJSON()
		return
	}
	u.Data["json"] = map[string]string{"uid": fmt.Sprint(uid)}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title GetPaginateObjectOfUser
// @Description find object by objectid
// @Param	userID		query 	int		true		"the userID you want to get"
// @Param	page		query	int		true		"the page"
// @Param	count		query	int		true		"the page length"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /page-object-users [get]
func (o *UserController) GetPaginateObjectOfUser() {
	userID, _ := o.GetInt32("userID")
	page, _ := o.GetInt32("page")
	count, _ := o.GetInt32("count")
	ob, err := models.GetPaginateObjectInUser(userID, page, count)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = ob
	}
	o.ServeJSON()
}



// @Title GetPaginateUser
// @Description find object by objectid
// @Param	page		query	int		true		"the page"
// @Param	count		query	int		true		"the page length"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /page-users [get]
func (o *UserController) GetPaginateUser() {
	page, _ := o.GetInt32("page")
	count, _ := o.GetInt32("count")
	ob, err := models.GetPaginateUsers(page, count)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = ob
	}
	o.ServeJSON()
}