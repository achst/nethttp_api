package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hopehook/nethttp_api/handler"
	"github.com/julienschmidt/httprouter"
)

func WebsocketHandler(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	switch r.Method {
	case "GET":
		switch Params.ByName("action") {
		case "upgrade":
			upgrade(w, r)
		default:
			handler.RouteErr(w)
		}
	case "POST":
		switch Params.ByName("action") {

		default:
			handler.RouteErr(w)
		}
	default:
		handler.RouteErr(w)
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(w, r, http.Header{})
	if err != nil {
		Logger.Error(err.Error())
		return
	}
	if err := hijackHandler(ws); err != nil {
		Logger.Error(err.Error())
		return
	}
	return
}

// hijackHandler is called on hijacked connection.
var hijackHandler = func(c *websocket.Conn) error {
	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			return err
		}
		if err = c.WriteMessage(messageType, p); err != nil {
			return err
		}
	}
	return nil
}
