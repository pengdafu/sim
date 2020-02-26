package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WsHandler(c *gin.Context)  {
	conn, err := Upgrader(c.Writer, c.Request)
	if err != nil {
		fmt.Printf("连接失败, %v", err)
		c.JSON(500, gin.H{
			"msg": "无法连接",
			"errmsg": err,
		})
	}
	go wsHandler(conn)
	select {}
}

func wsHandler(conn *websocket.Conn) {
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(msg))
		_ = conn.WriteMessage(t, msg)
	}
}