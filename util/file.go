package util

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(file multipart.File, path string) error {
	// 打开目标地址，把内容存进去
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if err != nil {
		return err
	}

	if _, err = io.Copy(f, file); err != nil {
		return err
	}
	return nil
}

// 使用接口检查是否有 Size() 方法
type fileSizer interface {
	Size() int64
}

// 从 multipart.File 获取文件大小
func GetUploadFileSize(f multipart.File) (int64, error) {
	// 从内存读取出来
	// if return *http.sectionReader, it is alias to *io.SectionReader
	if s, ok := f.(fileSizer); ok {
		return s.Size(), nil
	}
	// 从临时文件读取出来
	// or *os.File
	if fp, ok := f.(*os.File); ok {
		fi, err := fp.Stat()
		if err != nil {
			return 0, err
		}
		return fi.Size(), nil
	}
	return 0, nil
}
