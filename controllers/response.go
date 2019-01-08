package controllers

import (
	"github.com/astaxie/beego"
)

type ResponseController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Errcode int         `json:"code"`
	Errmsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
	Pager   interface{} `json:"page,omitempty"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"code"`
	Errmsg  interface{} `json:"msg"`
}
