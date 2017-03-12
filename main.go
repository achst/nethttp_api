// router: github.com/julienschmidt/httprouter
// mysql: github.com/go-sql-driver/mysql
// redis: github.com/garyburd/redigo/redis
// logs: github.com/astaxie/beego/logs
// websocket: github.com/gorilla/websocket
package main

import (
	"fmt"

	"github.com/hopehook/nethttp_api/config"
	"github.com/hopehook/nethttp_api/lib"
	"github.com/hopehook/nethttp_api/router"

	"net/http"
)

func main() {
	host := fmt.Sprintf("%s:%s", config.DEFAULT_SVR["ip"], config.DEFAULT_SVR["port"])
	lib.Logger.Info("System is running on %s", host)
	if err := http.ListenAndServe(host, router.Router); err != nil {
		lib.Logger.Error("Start nethttp failed:", err.Error())
	}
}
