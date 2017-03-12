package router

import (
	h "github.com/hopehook/nethttp_api/handler"
	"github.com/hopehook/nethttp_api/handler/template"
	"github.com/hopehook/nethttp_api/handler/websocket"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router = httprouter.New()

func init() {
	// static目录
	Router.NotFound = http.FileServer(http.Dir("static"))
	// login
	Router.POST("/login/", h.Raw(h.LoginHandler))
	Router.GET("/logout/", h.Raw(h.LogoutHandler))
	// file目录
	Router.GET("/tool/:action/", h.Raw(h.ToolHandler))
	Router.POST("/tool/:action/", h.Auth(h.ToolHandler))
	// template目录
	Router.GET("/t/:action/", h.Raw(template.TemplateHandler))
	Router.POST("/t/:action/", h.Raw(template.TemplateHandler))
	// websocket
	Router.GET("/websocket/:action/", h.Raw(websocket.WebsocketHandler)) // 建立websocket

	// 业务测试
	Router.GET("/user/:action/", h.Raw(h.UserHandler))
	Router.POST("/user/:action/", h.Raw(h.UserHandler))

}
