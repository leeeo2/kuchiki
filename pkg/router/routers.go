package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leeexeo/kuchiki/pkg/common/middleware"
	"github.com/leeexeo/kuchiki/pkg/common/setting"
	"github.com/leeexeo/kuchiki/pkg/controller"
)

func Setup() error {
	router := gin.Default()
	router.Use(middleware.RespHandleMiddleWare())

	router.POST("/user/add", controller.AddUser)
	router.POST("/user/delete", controller.DeleteUser)

	router.POST("/article/add", controller.AddArticle)

	serverConf := setting.GlobalConfig().Server
	listen := serverConf.ListenAddr + ":" + strconv.Itoa(serverConf.ListenPort)
	router.Run(listen)
	return nil
}
