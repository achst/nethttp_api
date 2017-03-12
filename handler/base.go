package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/hopehook/nethttp_api/define"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func loginAuth(w http.ResponseWriter, r *http.Request) (ok bool) {
	// 获取cookie的uid,sid
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		return false
	}
	sidCookie, err := r.Cookie("sid")
	if err != nil {
		return false
	}
	var uid = uidCookie.Value
	var sid = sidCookie.Value
	// 获取sid对应的session
	session_info, err := Cache.GetHashMapString(sid)
	if err != nil {
		return false
		Logger.Error(err.Error())
	}
	// cookie和session信息比对
	if session_info["uid"] != uid {
		return false
	}
	Logger.Info("auth success")
	return true
}

// login auth handler
func Auth(h httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
		printBeforeLog(w, r)
		hasAuth := loginAuth(w, r)

		if hasAuth {
			// Delegate request to the given handle
			h(w, r, Params)
			return
		}
		// Request Basic Authentication otherwise
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	})
}

// raw handler
func Raw(h httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
		printBeforeLog(w, r)
		h(w, r, Params)
		return
	})
}

// Write to response body
func WriteString(w http.ResponseWriter, result string) {
	printFinishLog(result)
	w.Write([]byte(result))
}

func WriteBytes(w http.ResponseWriter, result []byte) {
	printFinishLog(string(result))
	w.Write(result)
}

func CommonWriteSuccess(w http.ResponseWriter) {
	result := fmt.Sprintf(`{"ret": %d, "msg": %s, "data": {}}`, define.SUCCESS, define.SUCCESS_MSG)
	WriteString(w, result)
}

func CommonWriteError(w http.ResponseWriter) {
	result := fmt.Sprintf(`{"ret": %d, "msg": %s, "data": {}}`, define.ERR, define.ERR_MSG)
	WriteString(w, result)
}

func CommonWrite(w http.ResponseWriter, ret int64, msg string, data interface{}) {
	resultMap := map[string]interface{}{
		"ret":  ret,
		"msg":  msg,
		"data": data,
	}
	result, _ := json.Marshal(resultMap)
	WriteBytes(w, result)
}

func Render(w http.ResponseWriter, path string) {
	t, _ := template.ParseFiles(filepath.Join(TPL_PATH, path))
	t.Execute(w, nil)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8") // 渲染
}

// 打印handler执行之前的日志
func printBeforeLog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Logger.Info("================================================")
	Logger.Info("收到请求: method: %s, url: %s, from: %s", r.Method, r.URL.Path, r.RemoteAddr)
	Logger.Info("上传参数: all: %v, post: %v", r.Form, r.PostForm)
}

// 打印handler执行结束的日志
func printFinishLog(result interface{}) {
	Logger.Info("请求返回: %v ", result)
	Logger.Info("================================================")
}

func RouteErr(w http.ResponseWriter) {
	http.Error(w, "Unsupported path", http.StatusNotFound)
}
