package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.C.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}
