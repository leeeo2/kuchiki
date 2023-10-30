package controller

import "github.com/gin-gonic/gin"

// import (
// 	"reflect"

// 	"github.com/gin-gonic/gin"
// )

// type HandlerFunc func(ctx *gin.Context, input interface{}) (interface{}, error)

// type Controller struct {
// 	Group   string       // router group
// 	Path    string       // router path
// 	Method  string       // router method
// 	Req     reflect.Type // router request type
// 	Handler HandlerFunc  // router handler
// }

// var controllers = []*Controller{
// 	NewCreateUserController(),
// }

// func Setup() error {
// 	return nil
// }

// func RegisterRouter(router *gin.Engine, controlers []*Controller) error {
// 	for _, c := range controlers {
// 		router.Handle(c.Method, c.Path)
// 	}
// }

func Response(ctx *gin.Context, result interface{}, err error) {
	ctx.Set("result", result)
	ctx.Set("error", err)
}
