package handler

import (
	"github.com/hopehook/nethttp_api/define"

	"github.com/julienschmidt/httprouter"

	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	switch r.Method {
	case "GET":
		switch Params.ByName("action") {
		case "list":
			list(w, r)
		default:
			RouteErr(w)
		}

	case "POST":
		switch Params.ByName("action") {
		case "update":
			update(w, r)

		default:
			RouteErr(w)
		}
	default:
		RouteErr(w)
	}

}

func list(w http.ResponseWriter, r *http.Request) {
	sql := `select * from users`
	rows, _ := DB.Query(sql)
	CommonWrite(w, define.SUCCESS, define.SUCCESS_MSG, rows)
}

func update(w http.ResponseWriter, r *http.Request) {
	sql := `update users set is_deleted = 0 where id = 43`
	DB.Update(sql)
	CommonWriteSuccess(w)
}
