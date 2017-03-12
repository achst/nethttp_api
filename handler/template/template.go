package template

import (
	"github.com/hopehook/nethttp_api/handler"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func TemplateHandler(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	switch r.Method {
	case "GET":
		switch Params.ByName("action") {
		case "home":
			home(w, r)
		case "login":
			login(w, r)
		case "ws":
			ws(w, r)
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

func home(w http.ResponseWriter, r *http.Request) {
	handler.Render(w, "home.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	handler.Render(w, "login.html")
}

func ws(w http.ResponseWriter, r *http.Request) {
	handler.Render(w, "websocket.html")
}
