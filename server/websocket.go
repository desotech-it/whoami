package server

import (
	"fmt"
	"net/http"

	"github.com/desotech-it/whoami/app"
	"github.com/desotech-it/whoami/view"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	app.LogRequest(r)

	// If not a WebSocket upgrade request, serve the HTML test page
	if r.Header.Get("Upgrade") != "websocket" {
		v := view.NewWebSocketView("WebSocket Test")
		v.Write(w)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		reply := fmt.Sprintf("[%s] %s", app.Info.Hostname, string(p))
		if err := conn.WriteMessage(messageType, []byte(reply)); err != nil {
			return
		}
	}
}
