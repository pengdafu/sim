package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pengdafu/sim/connectivity/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("static/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/ws", controller.WsHandler)
	return r
}
