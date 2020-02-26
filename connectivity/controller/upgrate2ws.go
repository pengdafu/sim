package controller

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func Upgrader(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	wsupgrader := websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
	}
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}