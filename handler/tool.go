package handler

import (
	"path/filepath"

	"github.com/hopehook/nethttp_api/define"
	"github.com/hopehook/nethttp_api/util"
	"github.com/julienschmidt/httprouter"

	"mime/multipart"
	"net/http"
)

func ToolHandler(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	switch r.Method {
	case "GET":
		switch Params.ByName("action") {
		case "download":
			download(w, r)
		default:
			RouteErr(w)
		}

	case "POST":
		switch Params.ByName("action") {
		case "upload":
			upload(w, r)
		case "batch_upload":
			batchUpload(w, r)
		default:
			RouteErr(w)
		}
	default:
		RouteErr(w)
	}

}

// 下载"file/download/"目录下的文件
func download(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("filename")
	filepath := filepath.Join(DOWNLOAD_PATH, filename)
	Logger.Info(filepath)
	http.ServeFile(w, r, filepath)
}

// 上传文件到"file/upload/"目录
func upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file") // file 是上传表单域的名字
	defer file.Close()                          // 此时上传内容的 IO 已经打开，需要手动关闭！！
	if err != nil {
		Logger.Error("get upload file fail:", err)
		CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}

	// 文件名和路径
	filename := fileHeader.Filename
	filepath := filepath.Join(UPLOAD_PATH, filename)

	// 保存
	err = util.SaveFile(file, filepath)
	if err != nil {
		Logger.Error("upload file error: %s", err.Error())
		CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
		return
	}

	fileSize, _ := util.GetUploadFileSize(file)
	Logger.Info("文件大小: %d", fileSize)
	CommonWriteSuccess(w)
}

// 批量上传文件到"file/upload/"目录
func batchUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var (
		file multipart.File
		err  error
	)
	for _, fileHeader := range r.MultipartForm.File["file"] {
		if file, err = fileHeader.Open(); err != nil {
			Logger.Error("open upload file fail:", fileHeader.Filename, err)
			CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
			return
		}

		// 文件名和路径
		filename := fileHeader.Filename
		filepath := filepath.Join(UPLOAD_PATH, filename)

		// 保存
		err = util.SaveFile(file, filepath)
		if err != nil {
			Logger.Error("upload file error: %s", err.Error())
			CommonWrite(w, define.FAILED_ERR, define.FAILED_ERR_MSG, struct{}{})
			return
		}

		file = nil
		Logger.Info("save:" + fileHeader.Filename + " ")
	}

	CommonWriteSuccess(w)
}
