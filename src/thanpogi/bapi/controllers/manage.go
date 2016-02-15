package controllers

import (
	"fmt"
	"strconv"
	"encoding/json"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"thanpogi/common/models"
)

//Testing
type ManageController struct {
	beego.Controller
}

// @Title delete
// @Description delete
// @Param	id	path string	true "the objectid you want to get"
// @Success 200 
// @Failure 403 body is empty
// @router /delete/:id([0-9]+) [delete]
func (this *ManageController) Delete() {
	// convert the string value to an int
	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{}

	// Check if the article exists first
	if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist(); exist {
		if _, err := o.Delete(&models.Article{Id: articleId}); err == nil {
			this.Data["json"] = "Record Deleted. "
		} else {
			this.Data["json"] = "Record couldn't be deleted. Reason: " + err.Error()
		}
	} else {
		this.Data["json"] = "Record Doesn't exist."
	}
	
	this.ServeJSON()
}

// @Title update
// @Description update
// @Param	id	path string	true "the objectid you want to get"
// @Success 200 
// @Failure 403 body is empty
// @router /update/:id([0-9]+) [put]
func (this *ManageController) Update() {
	o := orm.NewOrm()
	o.Using("default")

	// convert the string value to an int
	if articleId, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err == nil {
		article := models.Article{Id: articleId}
		// attempt to load the record from the database
		if o.Read(&article) == nil {
			// set the Client and Url properties to arbitrary values
			article.Client = "Sitepoint"
			article.Url = "http://www.google.com"
			if _, err := o.Update(&article); err == nil {
				this.Data["json"] = "Record Was Updated."
			}
		} else {
			this.Data["json"] = "Record Was NOT Updated. Couldn't find article matching id"
		}
	} else {
		this.Data["json"] = "Record Was NOT Updated. Couldn't convert id from a string to a number. " + err.Error()
	}

	this.ServeJSON()
}

// @Title view
// @Description view
// @Success 200 
// @Failure 403 body is empty
// @router /view [get]
func (this *ManageController) View() {
	o := orm.NewOrm()
	o.Using("default")

	var articles []*models.Article
	num, err := o.QueryTable("articles").All(&articles)

	if err != orm.ErrNoRows && num > 0 {
		this.Data["json"] = articles
	}
	
	this.ServeJSON()
}

// @Title add
// @Description add
// @Param body body	models.Article true	"The object content"
// @Success 200 
// @Failure 403 body is empty
// @router /add [post]
func (this *ManageController) Add() {

	o := orm.NewOrm()
	o.Using("default")

	var article models.Article
    json.Unmarshal(this.Ctx.Input.RequestBody, &article)

	if err := this.ParseForm(&article); err != nil {
		this.Data["json"] = "Couldn't parse the form. Reason: " + err.Error()
	} else {
		valid := validation.Validation{}
		isValid, _ := valid.Valid(article)

		if this.Ctx.Input.Method() == "POST" {
			if !isValid {
				this.Data["json"] = "Form didn't validate."
			} else {
				searchArticle := models.Article{Name: article.Name}
				beego.Debug("Article name supplied:", article.Name)
				err = o.Read(&searchArticle)
				beego.Debug("Err:", err)

				if err == orm.ErrNoRows || err == orm.ErrMissPK {
					beego.Debug("No article found matching details supplied. Attempting to insert article: ", article)
					id, err := o.Insert(&article)
					if err == nil {
						msg := fmt.Sprintf("Article inserted with id:", id)
						beego.Debug(msg)
					} else {
						msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
						beego.Debug(msg)
					}
				} else {
					beego.Debug("Article found matching details supplied. Cannot insert")
				}
			}
		}
	}
	
	this.ServeJSON()
}
