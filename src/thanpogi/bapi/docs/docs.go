package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/manage","description":"Testing\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/manage":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/manage","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/delete/:id([0-9]+)","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete","parameters":[{"paramType":"path","name":"id","description":"\"the objectid you want to get\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":0,"message":"200","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/update/:id([0-9]+)","description":"","operations":[{"httpMethod":"PUT","nickname":"update","type":"","summary":"update","parameters":[{"paramType":"path","name":"id","description":"\"the objectid you want to get\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":0,"message":"200","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/view","description":"","operations":[{"httpMethod":"GET","nickname":"view","type":"","summary":"view","responseMessages":[{"code":0,"message":"200","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/add","description":"","operations":[{"httpMethod":"POST","nickname":"add","type":"","summary":"add","parameters":[{"paramType":"body","name":"body","description":"\"The object content\"","dataType":"Article","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":0,"message":"200","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
