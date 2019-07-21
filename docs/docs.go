// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-21 13:30:47.72711 +0800 CST m=+0.055580549

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "淘宝App 闲鱼App 相关Api 调用，仅供测试使用。 点击接口展开可以查看详情，点击\"Try it out\" 填写参数即可进行接口测试 更多交流 3394772548@qq.com",
        "title": "yue-spider API文档",
        "contact": {
            "name": "API Support",
            "email": "3394772548@qq.com"
        },
        "license": {},
        "version": "1.0"
    },
    "basePath": "/",
    "paths": {
        "/mtop.taobao.detail.getdesc": {
            "get": {
                "description": "获取淘宝天猫宝贝详情描述接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "淘宝商品信息"
                ],
                "summary": "淘宝商品宝贝详情描述接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "淘宝商品id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.detail.getdetail": {
            "get": {
                "description": "获取淘宝天猫宝贝基本信息接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "淘宝商品信息"
                ],
                "summary": "淘宝App宝贝基本信息接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "淘宝商品id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.idle.item.detail": {
            "get": {
                "description": "获取闲鱼宝贝基本信息接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "闲鱼商品信息"
                ],
                "summary": "闲鱼App宝贝基本信息接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "闲鱼商品id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.idle.main.item.search": {
            "get": {
                "description": "闲鱼搜索接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "闲鱼商品信息"
                ],
                "summary": "闲鱼搜索接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码(默认1)",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.rate.detaillist.get": {
            "get": {
                "description": "获取淘宝天猫宝贝宝贝评论接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "淘宝商品信息"
                ],
                "summary": "淘宝宝贝评论接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "淘宝商品id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码(默认1)",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.sharepassword.querypassword": {
            "get": {
                "description": "淘口令解析",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "淘宝客"
                ],
                "summary": "淘口令解析",
                "parameters": [
                    {
                        "type": "string",
                        "description": "淘口令",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.ugc.tql.facade": {
            "get": {
                "description": "获取淘宝 天猫 宝贝买家秀",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "淘宝商品信息"
                ],
                "summary": "淘宝App宝贝买家秀",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "淘宝商品id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/mtop.taobao.wsearch.appsearch": {
            "get": {
                "description": "淘宝搜索接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "淘宝商品信息"
                ],
                "summary": "淘宝App搜索接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页码(默认1)",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
