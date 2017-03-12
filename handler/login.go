package handler

import (
	"strconv"

	"github.com/hopehook/nethttp_api/define"
	"github.com/julienschmidt/httprouter"

	"net/http"
	"time"
)

// login
func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// upload param
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	Logger.Info(username)
	Logger.Info(password)
	// auth
	sql := `SELECT id AS uid, password, phone, name FROM users WHERE phone = ? AND is_deleted = 0 AND app_type = 1`
	rows, _ := DB.Query(sql, username)
	if len(rows) != 1 {
		CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	userInfo := rows[0]
	if password != userInfo["password"] {
		CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// session
	uid := userInfo["uid"].(int64)
	uidStr := strconv.Itoa(int(uid))
	sid := uidStr
	_, err := Cache.SetHashMap(sid, userInfo)
	if err != nil {
		CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// cookie
	c := &http.Cookie{
		Name:     "uid",
		Value:    uidStr,
		Expires:  time.Now().AddDate(1, 0, 7),
		MaxAge:   3600 * 7,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, c)
	c.Name = "sid"
	c.Value = sid
	http.SetCookie(w, c)

	// login success
	CommonWriteSuccess(w)
}

// logout
func LogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 获取cookie的uid,sid
	var sid = r.FormValue("sid")
	// session
	_, err := Cache.DelKey(sid)
	if err != nil {
		CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}
	// cookie
	c := &http.Cookie{
		Name:     "uid",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1, // equivalently 'Max-Age: 0'
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, c)
	c.Name = "sid"
	c.Value = ""
	http.SetCookie(w, c)

	// logout success
	CommonWriteSuccess(w)
}
